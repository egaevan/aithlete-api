package goal

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Goal, error)
	FindByUserID(ctx context.Context, userID string) ([]Goal, error)
	Create(ctx context.Context, goal *Goal) error
	Update(ctx context.Context, goal *Goal) error
	Delete(ctx context.Context, id string) error
}
