package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type WorkoutRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Workout, error)
	FindByUserID(ctx context.Context, userID string) ([]entity.Workout, error)
	Create(ctx context.Context, workout *entity.Workout) error
	Update(ctx context.Context, workout *entity.Workout) error
	Delete(ctx context.Context, id string) error
}
