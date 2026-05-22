package exercise

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type ListExercisesUseCase interface {
	List(ctx context.Context) ([]dto.ExerciseDetailResult, error)
}

type listExercisesUseCase struct {
	exerciseRepo repository.ExerciseRepository
}

func NewListExercisesUseCase(exerciseRepo repository.ExerciseRepository) ListExercisesUseCase {
	return &listExercisesUseCase{exerciseRepo: exerciseRepo}
}

func (uc *listExercisesUseCase) List(ctx context.Context) ([]dto.ExerciseDetailResult, error) {
	exercises, err := uc.exerciseRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("list exercises: %w", err)
	}
	return mapper.ExerciseToResultList(exercises), nil
}
