package schedule

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Schedule, error)
	FindByUserID(ctx context.Context, userID string) ([]Schedule, error)
	FindByUserIDAndDate(ctx context.Context, userID, date string) ([]Schedule, error)
	Create(ctx context.Context, schedule *Schedule) error
	Update(ctx context.Context, schedule *Schedule) error
	Delete(ctx context.Context, id string) error
}
