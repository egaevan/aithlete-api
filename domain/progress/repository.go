package progress

import "context"

type Repository interface {
	FindBodyWeightByUserID(ctx context.Context, userID string) ([]BodyWeight, error)
	AddBodyWeight(ctx context.Context, bw *BodyWeight) error
	FindStrengthByUserID(ctx context.Context, userID string) ([]StrengthRecord, error)
	FindConsistency(ctx context.Context, userID string) ([]Consistency, error)
	FindMuscleVolume(ctx context.Context, userID string) ([]MuscleVolume, error)
}
