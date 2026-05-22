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

type CompleteWorkoutUseCase interface {
	Complete(ctx context.Context, userID, workoutID string) (*dto.WorkoutResult, error)
}

type completeWorkoutUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewCompleteWorkoutUseCase(workoutRepo repository.WorkoutRepository) CompleteWorkoutUseCase {
	return &completeWorkoutUseCase{workoutRepo: workoutRepo}
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

	return mapper.WorkoutToResult(w), nil
}
