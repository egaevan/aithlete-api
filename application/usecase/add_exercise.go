package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type AddExerciseUseCase interface {
	AddExercise(ctx context.Context, userID, workoutID string, exercise entity.ExerciseRef, sets []entity.Set) (*dto.WorkoutResult, error)
}

type addExerciseUseCase struct {
	workoutRepo repository.WorkoutRepository
}

func NewAddExerciseUseCase(workoutRepo repository.WorkoutRepository) AddExerciseUseCase {
	return &addExerciseUseCase{workoutRepo: workoutRepo}
}

func (uc *addExerciseUseCase) AddExercise(ctx context.Context, userID, workoutID string, exercise entity.ExerciseRef, sets []entity.Set) (*dto.WorkoutResult, error) {
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

	we := entity.WorkoutExercise{
		Exercise: exercise,
		Sets:     sets,
	}

	if err := w.AddExercise(we); err != nil {
		return nil, err
	}

	if err := uc.workoutRepo.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("update workout: %w", err)
	}

	return mapper.WorkoutToResult(w), nil
}
