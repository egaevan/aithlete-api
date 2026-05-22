package analytics

import analyticsuc "github.com/aithlete/aithlete-api/application/usecase/analytics"

type Handler struct {
	dashboardUseCase              analyticsuc.GetDashboardUseCase
	weeklyProgressUseCase         analyticsuc.GetWeeklyProgressUseCase
	streakUseCase                 analyticsuc.GetStreakUseCase
	overviewUseCase               analyticsuc.GetOverviewUseCase
	weeklyVolumeUseCase           analyticsuc.GetWeeklyVolumeUseCase
	muscleVolumeDistributionUseCase analyticsuc.GetMuscleVolumeDistributionUseCase
}

func New(
	dashboard analyticsuc.GetDashboardUseCase,
	weeklyProgress analyticsuc.GetWeeklyProgressUseCase,
	streak analyticsuc.GetStreakUseCase,
	overview analyticsuc.GetOverviewUseCase,
	weeklyVolume analyticsuc.GetWeeklyVolumeUseCase,
	muscleVolumeDistribution analyticsuc.GetMuscleVolumeDistributionUseCase,
) *Handler {
	return &Handler{
		dashboardUseCase:              dashboard,
		weeklyProgressUseCase:         weeklyProgress,
		streakUseCase:                 streak,
		overviewUseCase:               overview,
		weeklyVolumeUseCase:           weeklyVolume,
		muscleVolumeDistributionUseCase: muscleVolumeDistribution,
	}
}
