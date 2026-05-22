package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type DeleteWorkoutUseCase interface {
	Delete(ctx context.Context, userID, workoutID string) error
}

type deleteWorkoutUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewDeleteWorkoutUseCase(workoutRepo repository.WorkoutRepository) DeleteWorkoutUseCase {
	return &deleteWorkoutUseCase{workoutRepo: workoutRepo}
}

func (uc *deleteWorkoutUseCase) Delete(ctx context.Context, userID, workoutID string) error {
	w, err := uc.workoutRepo.FindByID(ctx, workoutID)
	if err != nil {
		if service.IsWorkoutNotFound(err) {
			return domainerr.ErrWorkoutNotFound
		}
		return fmt.Errorf("find workout: %w", err)
	}

	if w.UserID != userID {
		return domainerr.ErrWorkoutNotFound
	}

	if err := uc.workoutRepo.Delete(ctx, workoutID); err != nil {
		return fmt.Errorf("delete workout: %w", err)
	}

	return nil
}
