package workout

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type GetWorkoutUseCase interface {
	Get(ctx context.Context, userID, workoutID string) (*dto.WorkoutResult, error)
}

type getWorkoutUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewGetWorkoutUseCase(workoutRepo repository.WorkoutRepository) GetWorkoutUseCase {
	return &getWorkoutUseCase{workoutRepo: workoutRepo}
}

func (uc *getWorkoutUseCase) Get(ctx context.Context, userID, workoutID string) (*dto.WorkoutResult, error) {
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

	return mapper.WorkoutToResult(w), nil
}
