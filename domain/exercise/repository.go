package exercise

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Exercise, error)
	FindByID(ctx context.Context, id string) (*Exercise, error)
	FindByMuscleGroup(ctx context.Context, muscleGroup string) ([]Exercise, error)
	Create(ctx context.Context, exercise *Exercise) error
}
