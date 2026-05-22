package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type ListWorkoutsUseCase interface {
	List(ctx context.Context, userID string) ([]dto.WorkoutResult, error)
}

type listWorkoutsUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewListWorkoutsUseCase(workoutRepo repository.WorkoutRepository) ListWorkoutsUseCase {
	return &listWorkoutsUseCase{workoutRepo: workoutRepo}
}

func (uc *listWorkoutsUseCase) List(ctx context.Context, userID string) ([]dto.WorkoutResult, error) {
	workouts, err := uc.workoutRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list workouts: %w", err)
	}

	results := make([]dto.WorkoutResult, len(workouts))
	for i := range workouts {
		results[i] = *mapper.WorkoutToResult(&workouts[i])
	}

	return results, nil
}
