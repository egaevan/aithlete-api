package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func DashboardToResult(d *entity.Dashboard) *dto.DashboardResult {
	return &dto.DashboardResult{
		Stats:            DashboardStatsToResult(&d.Stats),
		WeeklyProgress:   WeeklyProgressDayToResultList(d.WeeklyProgress),
		Streak:           StreakToResult(&d.Streak),
		RecentActivity:   ActivityToResultList(d.RecentActivity),
		MuscleRecovery:   MuscleRecoveryToResultList(d.MuscleRecovery),
		TodaySchedule:    TodayScheduleItemToResultList(d.TodaySchedule),
		AIRecommendation: AIRecommendationSummaryToResult(d.AIRecommendation),
	}
}

func DashboardStatsToResult(s *entity.DashboardStats) dto.DashboardStatsResult {
	return dto.DashboardStatsResult{
		CaloriesBurned:     s.CaloriesBurned,
		CaloriesTrend:      s.CaloriesTrend,
		ActiveMinutes:      s.ActiveMinutes,
		ActiveMinutesTrend: s.ActiveMinutesTrend,
		GoalsCompleted:     s.GoalsCompleted,
		GoalsTotal:         s.GoalsTotal,
		GoalsTrend:         s.GoalsTrend,
		AvgHeartRate:       s.AvgHeartRate,
		HeartRateTrend:     s.HeartRateTrend,
	}
}

func WeeklyProgressDayToResult(w *entity.WeeklyProgressDay) dto.WeeklyProgressDayResult {
	return dto.WeeklyProgressDayResult{
		Day:      w.Day,
		Calories: w.Calories,
		Duration: w.Duration,
	}
}

func WeeklyProgressDayToResultList(ws []entity.WeeklyProgressDay) []dto.WeeklyProgressDayResult {
	result := make([]dto.WeeklyProgressDayResult, len(ws))
	for i := range ws {
		result[i] = WeeklyProgressDayToResult(&ws[i])
	}
	return result
}

func StreakToResult(s *entity.Streak) dto.StreakResult {
	return dto.StreakResult{
		Current:      s.Current,
		PersonalBest: s.PersonalBest,
		History:      s.History,
	}
}

func ActivityToResult(a *entity.Activity) dto.ActivityResult {
	return dto.ActivityResult{
		ID:          a.ID,
		Type:        a.Type,
		Title:       a.Title,
		Description: a.Description,
		Timestamp:   a.Timestamp.Format(time.RFC3339),
	}
}

func ActivityToResultList(as []entity.Activity) []dto.ActivityResult {
	result := make([]dto.ActivityResult, len(as))
	for i := range as {
		result[i] = ActivityToResult(&as[i])
	}
	return result
}

func MuscleRecoveryToResult(m *entity.MuscleRecovery) dto.MuscleRecoveryResult {
	return dto.MuscleRecoveryResult{
		MuscleGroup:      m.MuscleGroup,
		Recovery:         m.Recovery,
		ReadyForTraining: m.ReadyForTraining,
	}
}

func MuscleRecoveryToResultList(ms []entity.MuscleRecovery) []dto.MuscleRecoveryResult {
	result := make([]dto.MuscleRecoveryResult, len(ms))
	for i := range ms {
		result[i] = MuscleRecoveryToResult(&ms[i])
	}
	return result
}

func TodayScheduleItemToResult(t *entity.TodayScheduleItem) dto.TodayScheduleItemResult {
	return dto.TodayScheduleItemResult{
		ID:        t.ID,
		Time:      t.Time,
		Title:     t.Title,
		Duration:  t.Duration,
		Type:      t.Type,
		Completed: t.Completed,
	}
}

func TodayScheduleItemToResultList(ts []entity.TodayScheduleItem) []dto.TodayScheduleItemResult {
	result := make([]dto.TodayScheduleItemResult, len(ts))
	for i := range ts {
		result[i] = TodayScheduleItemToResult(&ts[i])
	}
	return result
}

func AIRecommendationSummaryToResult(a *entity.AIRecommendationSummary) *dto.AIRecommendationSummaryResult {
	if a == nil {
		return nil
	}
	return &dto.AIRecommendationSummaryResult{
		ID:          a.ID,
		Type:        a.Type,
		Title:       a.Title,
		Description: a.Description,
		Confidence:  a.Confidence,
	}
}

func WeeklyVolumeToResult(w *entity.WeeklyVolume) dto.WeeklyVolumeResult {
	return dto.WeeklyVolumeResult{
		Week:   w.Week,
		Volume: w.Volume,
	}
}

func WeeklyVolumeToResultList(ws []entity.WeeklyVolume) []dto.WeeklyVolumeResult {
	result := make([]dto.WeeklyVolumeResult, len(ws))
	for i := range ws {
		result[i] = WeeklyVolumeToResult(&ws[i])
	}
	return result
}

func MuscleVolumeDistributionToResult(m *entity.MuscleVolumeDistribution) dto.MuscleVolumeDistributionResult {
	return dto.MuscleVolumeDistributionResult{
		Muscle: m.Muscle,
		Volume: m.Volume,
	}
}

func AnalyticsOverviewToResult(o *entity.AnalyticsOverview) *dto.AnalyticsOverviewResult {
	return &dto.AnalyticsOverviewResult{
		TotalVolume:          o.TotalVolume,
		TotalVolumeTrend:     o.TotalVolumeTrend,
		AvgSession:           o.AvgSession,
		AvgSessionTrend:      o.AvgSessionTrend,
		SessionsPerMonth:     o.SessionsPerMonth,
		SessionsPerMonthTrend: o.SessionsPerMonthTrend,
		GoalCompletion:       o.GoalCompletion,
		GoalCompletionTrend:  o.GoalCompletionTrend,
	}
}

func MuscleVolumeDistributionToResultList(ms []entity.MuscleVolumeDistribution) []dto.MuscleVolumeDistributionResult {
	result := make([]dto.MuscleVolumeDistributionResult, len(ms))
	for i := range ms {
		result[i] = MuscleVolumeDistributionToResult(&ms[i])
	}
	return result
}
