package analytics

import "context"

type Repository interface {
	GetDashboard(ctx context.Context, userID string) (*Dashboard, error)
	GetWeeklyProgress(ctx context.Context, userID string) ([]WeeklyProgressDay, error)
	GetStreak(ctx context.Context, userID string) (*Streak, error)
	GetWeeklyVolume(ctx context.Context, userID string) ([]WeeklyVolume, error)
	GetMuscleVolumeDistribution(ctx context.Context, userID string) ([]MuscleVolumeDistribution, error)
}
