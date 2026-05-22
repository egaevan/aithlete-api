package workout

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type CreateWorkoutUseCase interface {
	Create(ctx context.Context, userID, name, date, weightUnit, notes string) (*dto.WorkoutResult, error)
}

type createWorkoutUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewCreateWorkoutUseCase(workoutRepo repository.WorkoutRepository) CreateWorkoutUseCase {
	return &createWorkoutUseCase{workoutRepo: workoutRepo}
}

func (uc *createWorkoutUseCase) Create(ctx context.Context, userID, name, date, weightUnit, notes string) (*dto.WorkoutResult, error) {
	w := entity.NewWorkout(userID, name, date, weightUnit, notes)
	if err := uc.workoutRepo.Create(ctx, w); err != nil {
		return nil, fmt.Errorf("create workout: %w", err)
	}

	return mapper.WorkoutToResult(w), nil
}
