package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type AnalyticsRepository interface {
	GetDashboard(ctx context.Context, userID string) (*entity.Dashboard, error)
	GetWeeklyProgress(ctx context.Context, userID string) ([]entity.WeeklyProgressDay, error)
	GetStreak(ctx context.Context, userID string) (*entity.Streak, error)
	GetWeeklyVolume(ctx context.Context, userID string) ([]entity.WeeklyVolume, error)
	GetMuscleVolumeDistribution(ctx context.Context, userID string) ([]entity.MuscleVolumeDistribution, error)
	GetOverview(ctx context.Context, userID string) (*entity.AnalyticsOverview, error)
}
