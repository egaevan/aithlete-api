package workout

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Workout, error)
	FindByUserID(ctx context.Context, userID string) ([]Workout, error)
	Create(ctx context.Context, workout *Workout) error
	Update(ctx context.Context, workout *Workout) error
	Delete(ctx context.Context, id string) error
}
