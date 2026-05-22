package exercise

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type GetExerciseUseCase interface {
	Get(ctx context.Context, id string) (*dto.ExerciseDetailResult, error)
}

type getExerciseUseCase struct {
	exerciseRepo repository.ExerciseRepository
}

func NewGetExerciseUseCase(exerciseRepo repository.ExerciseRepository) GetExerciseUseCase {
	return &getExerciseUseCase{exerciseRepo: exerciseRepo}
}

func (uc *getExerciseUseCase) Get(ctx context.Context, id string) (*dto.ExerciseDetailResult, error) {
	e, err := uc.exerciseRepo.FindByID(ctx, id)
	if err != nil {
		if service.IsExerciseNotFound(err) {
			return nil, domainerr.ErrExerciseNotFound
		}
		return nil, fmt.Errorf("find exercise: %w", err)
	}
	return mapper.ExerciseToResult(e), nil
}
