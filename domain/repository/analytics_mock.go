package repository

import (
	"context"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockAnalyticsRepository struct{}

func NewMockAnalyticsRepository() *MockAnalyticsRepository {
	return &MockAnalyticsRepository{}
}

func (m *MockAnalyticsRepository) GetDashboard(_ context.Context, userID string) (*entity.Dashboard, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	now := time.Now()
	return &entity.Dashboard{
		Stats: entity.DashboardStats{
			CaloriesBurned:     2450,
			CaloriesTrend:      "+12%",
			ActiveMinutes:      75,
			ActiveMinutesTrend: "+5%",
			GoalsCompleted:     3,
			GoalsTotal:         5,
			GoalsTrend:         "on track",
			AvgHeartRate:       138,
			HeartRateTrend:     "-2%",
		},
		WeeklyProgress: []entity.WeeklyProgressDay{
			{Day: "Mon", Calories: 450, Duration: 45},
			{Day: "Tue", Calories: 320, Duration: 30},
			{Day: "Wed", Calories: 520, Duration: 55},
			{Day: "Thu", Calories: 0, Duration: 0},
			{Day: "Fri", Calories: 480, Duration: 50},
			{Day: "Sat", Calories: 600, Duration: 60},
			{Day: "Sun", Calories: 0, Duration: 0},
		},
		Streak: entity.Streak{
			Current:      5,
			PersonalBest: 14,
			History:      []bool{true, true, true, true, true, false, false},
		},
		RecentActivity: []entity.Activity{
			{ID: "act-1", Type: "workout", Title: "Morning Run", Description: "5km run", Timestamp: now.Add(-2 * time.Hour)},
			{ID: "act-2", Type: "workout", Title: "Upper Body", Description: "Chest and triceps", Timestamp: now.Add(-26 * time.Hour)},
		},
		MuscleRecovery: []entity.MuscleRecovery{
			{MuscleGroup: "chest", Recovery: 85, ReadyForTraining: true},
			{MuscleGroup: "legs", Recovery: 45, ReadyForTraining: false},
		},
		TodaySchedule: []entity.TodayScheduleItem{
			{ID: "sched-1", Time: "06:30", Title: "Morning Run", Duration: "30 min", Type: "cardio", Completed: true},
		},
	}, nil
}

func (m *MockAnalyticsRepository) GetWeeklyProgress(_ context.Context, userID string) ([]entity.WeeklyProgressDay, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return []entity.WeeklyProgressDay{
		{Day: "Mon", Calories: 450, Duration: 45},
		{Day: "Tue", Calories: 320, Duration: 30},
		{Day: "Wed", Calories: 520, Duration: 55},
	}, nil
}

func (m *MockAnalyticsRepository) GetStreak(_ context.Context, userID string) (*entity.Streak, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return &entity.Streak{
		Current:      5,
		PersonalBest: 14,
		History:      []bool{true, true, true, true, true, false, false},
	}, nil
}

func (m *MockAnalyticsRepository) GetWeeklyVolume(_ context.Context, userID string) ([]entity.WeeklyVolume, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return []entity.WeeklyVolume{
		{Week: "2026-W19", Volume: 12500},
		{Week: "2026-W20", Volume: 15800},
		{Week: "2026-W21", Volume: 14200},
	}, nil
}

func (m *MockAnalyticsRepository) GetMuscleVolumeDistribution(_ context.Context, userID string) ([]entity.MuscleVolumeDistribution, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return []entity.MuscleVolumeDistribution{
		{Muscle: "chest", Volume: 4500},
		{Muscle: "back", Volume: 5200},
		{Muscle: "legs", Volume: 6100},
	}, nil
}

func (m *MockAnalyticsRepository) GetOverview(_ context.Context, userID string) (*entity.AnalyticsOverview, error) {
	if userID == "" {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return &entity.AnalyticsOverview{
		TotalVolume:          80100,
		TotalVolumeTrend:     "+15%",
		AvgSession:           58,
		AvgSessionTrend:      "+5 min",
		SessionsPerMonth:     22,
		SessionsPerMonthTrend: "+3",
		GoalCompletion:       20,
		GoalCompletionTrend:  "+4%",
	}, nil
}

var _ AnalyticsRepository = (*MockAnalyticsRepository)(nil)
