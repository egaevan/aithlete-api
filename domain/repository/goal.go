package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type GoalRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Goal, error)
	FindByUserID(ctx context.Context, userID string) ([]entity.Goal, error)
	Create(ctx context.Context, goal *entity.Goal) error
	Update(ctx context.Context, goal *entity.Goal) error
	Delete(ctx context.Context, id string) error
}
