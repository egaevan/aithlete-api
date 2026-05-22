package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type ProgressRepository interface {
	FindBodyWeightByUserID(ctx context.Context, userID string) ([]entity.BodyWeight, error)
	AddBodyWeight(ctx context.Context, bw *entity.BodyWeight) error
	FindStrengthByUserID(ctx context.Context, userID string) ([]entity.StrengthRecord, error)
	AddStrengthRecord(ctx context.Context, sr *entity.StrengthRecord) error
	FindConsistency(ctx context.Context, userID string) ([]entity.Consistency, error)
	UpsertConsistency(ctx context.Context, userID, week string, completed int) error
	FindMuscleVolume(ctx context.Context, userID string) ([]entity.MuscleVolume, error)
	UpsertMuscleVolume(ctx context.Context, userID, muscleGroup string, volume float64) error
}
