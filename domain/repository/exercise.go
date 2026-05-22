package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type ExerciseRepository interface {
	FindAll(ctx context.Context) ([]entity.Exercise, error)
	FindByID(ctx context.Context, id string) (*entity.Exercise, error)
	FindByMuscleGroup(ctx context.Context, muscleGroup string) ([]entity.Exercise, error)
	Create(ctx context.Context, exercise *entity.Exercise) error
}
