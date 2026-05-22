package workout

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type CompleteWorkoutUseCase interface {
	Complete(ctx context.Context, userID, workoutID string) (*dto.WorkoutResult, error)
}

type completeWorkoutUseCase struct {
	workoutRepo  repository.WorkoutRepository
	progressRepo repository.ProgressRepository
}

func NewCompleteWorkoutUseCase(workoutRepo repository.WorkoutRepository, progressRepo repository.ProgressRepository) CompleteWorkoutUseCase {
	return &completeWorkoutUseCase{workoutRepo: workoutRepo, progressRepo: progressRepo}
}

func (uc *completeWorkoutUseCase) Complete(ctx context.Context, userID, workoutID string) (*dto.WorkoutResult, error) {
	w, err := uc.workoutRepo.FindByID(ctx, workoutID)
	if err != nil {
		if service.IsWorkoutNotFound(err) {
			return nil, domainerr.ErrWorkoutNotFound
		}
		return nil, fmt.Errorf("find workout: %w", err)
	}

	if w.UserID != userID {
		return nil, domainerr.ErrWorkoutNotFound
	}

	if err := w.Complete(); err != nil {
		return nil, err
	}

	if err := uc.workoutRepo.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("update workout: %w", err)
	}

	if err := uc.syncProgress(ctx, w); err != nil {
		return nil, fmt.Errorf("sync progress: %w", err)
	}

	return mapper.WorkoutToResult(w), nil
}

func (uc *completeWorkoutUseCase) syncProgress(ctx context.Context, w *entity.Workout) error {
	for _, we := range w.Exercises {
		volume := exerciseVolume(we.Sets)
		oneRepMax := estimateOneRepMax(we.Sets)

		sr := &entity.StrengthRecord{
			UserID:    w.UserID,
			Exercise:  we.Exercise.Name,
			Date:      w.Date,
			OneRepMax: oneRepMax,
			Volume:    volume,
		}
		if err := uc.progressRepo.AddStrengthRecord(ctx, sr); err != nil {
			return fmt.Errorf("add strength record: %w", err)
		}

		if we.Exercise.MuscleGroup != "" {
			if err := uc.progressRepo.UpsertMuscleVolume(ctx, w.UserID, we.Exercise.MuscleGroup, volume); err != nil {
				return fmt.Errorf("upsert muscle volume: %w", err)
			}
		}
	}

	t, err := time.Parse("2006-01-02", w.Date)
	if err != nil {
		t = time.Now()
	}
	year, week := t.ISOWeek()
	weekStr := fmt.Sprintf("%d-W%02d", year, week)

	if err := uc.progressRepo.UpsertConsistency(ctx, w.UserID, weekStr, 1); err != nil {
		return fmt.Errorf("upsert consistency: %w", err)
	}

	return nil
}

func exerciseVolume(sets []entity.Set) float64 {
	var total float64
	for _, s := range sets {
		total += float64(s.Reps) * s.Weight
	}
	return total
}

func estimateOneRepMax(sets []entity.Set) float64 {
	var maxWeight float64
	var bestReps int
	for _, s := range sets {
		if s.Weight > maxWeight {
			maxWeight = s.Weight
			bestReps = s.Reps
		}
	}
	if maxWeight == 0 {
		return 0
	}
	if bestReps == 1 {
		return maxWeight
	}
	return maxWeight * (1 + float64(bestReps)/30.0)
}
