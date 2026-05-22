package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type UpdateSetUseCase interface {
	UpdateSet(ctx context.Context, userID, workoutID, exerciseID, setID string, reps int, weight float64, rpe int) (*dto.WorkoutResult, error)
}

type updateSetUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewUpdateSetUseCase(workoutRepo repository.WorkoutRepository) UpdateSetUseCase {
	return &updateSetUseCase{workoutRepo: workoutRepo}
}

func (uc *updateSetUseCase) UpdateSet(ctx context.Context, userID, workoutID, exerciseID, setID string, reps int, weight float64, rpe int) (*dto.WorkoutResult, error) {
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

	if err := w.UpdateSet(exerciseID, setID, reps, weight, rpe); err != nil {
		return nil, err
	}

	if err := uc.workoutRepo.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("update workout: %w", err)
	}

	return mapper.WorkoutToResult(w), nil
}
