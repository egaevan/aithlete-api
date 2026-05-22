package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type ScheduleRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Schedule, error)
	FindByUserID(ctx context.Context, userID string) ([]entity.Schedule, error)
	FindByUserIDAndDate(ctx context.Context, userID, date string) ([]entity.Schedule, error)
	Create(ctx context.Context, schedule *entity.Schedule) error
	Update(ctx context.Context, schedule *entity.Schedule) error
	Delete(ctx context.Context, id string) error
}
