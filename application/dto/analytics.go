package dto

type DashboardResult struct {
	Stats            DashboardStatsResult            `json:"stats"`
	WeeklyProgress   []WeeklyProgressDayResult       `json:"weeklyProgress"`
	Streak           StreakResult                    `json:"streak"`
	RecentActivity   []ActivityResult                `json:"recentActivity"`
	MuscleRecovery   []MuscleRecoveryResult          `json:"muscleRecovery"`
	TodaySchedule    []TodayScheduleItemResult       `json:"todaySchedule"`
	AIRecommendation *AIRecommendationSummaryResult  `json:"aiRecommendation"`
}

type DashboardStatsResult struct {
	CaloriesBurned     int
	CaloriesTrend      string
	ActiveMinutes      int
	ActiveMinutesTrend string
	GoalsCompleted     int
	GoalsTotal         int
	GoalsTrend         string
	AvgHeartRate       int
	HeartRateTrend     string
}

type WeeklyProgressDayResult struct {
	Day      string
	Calories int
	Duration int
}

type StreakResult struct {
	Current      int
	PersonalBest int
	History      []bool
}

type ActivityResult struct {
	ID          string
	Type        string
	Title       string
	Description string
	Timestamp   string
}

type MuscleRecoveryResult struct {
	MuscleGroup      string
	Recovery         int
	ReadyForTraining bool
}

type TodayScheduleItemResult struct {
	ID        string
	Time      string
	Title     string
	Duration  string
	Type      string
	Completed bool
}

type AIRecommendationSummaryResult struct {
	ID          string
	Type        string
	Title       string
	Description string
	Confidence  float64
}

type AnalyticsOverviewResult struct {
	TotalVolume           int
	TotalVolumeTrend      string
	AvgSession            int
	AvgSessionTrend       string
	SessionsPerMonth      int
	SessionsPerMonthTrend string
	GoalCompletion        int
	GoalCompletionTrend   string
}

type WeeklyVolumeResult struct {
	Week   string
	Volume float64
}

type MuscleVolumeDistributionResult struct {
	Muscle string
	Volume float64
}
