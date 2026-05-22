package analytics

import "time"

type Dashboard struct {
	Stats          DashboardStats
	WeeklyProgress []WeeklyProgressDay
	Streak         Streak
	RecentActivity []Activity
	MuscleRecovery []MuscleRecovery
	TodaySchedule  []TodayScheduleItem
	AIRecommendation *AIRecommendationSummary
}

type DashboardStats struct {
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

type WeeklyProgressDay struct {
	Day      string
	Calories int
	Duration int
}

type Streak struct {
	Current      int
	PersonalBest int
	History      []bool
}

type Activity struct {
	ID          string
	Type        string
	Title       string
	Description string
	Timestamp   time.Time
}

type MuscleRecovery struct {
	MuscleGroup     string
	Recovery        int
	ReadyForTraining bool
}

type TodayScheduleItem struct {
	ID        string
	Time      string
	Title     string
	Duration  string
	Type      string
	Completed bool
}

type AIRecommendationSummary struct {
	ID          string
	Type        string
	Title       string
	Description string
	Confidence  float64
}

type WeeklyVolume struct {
	Week   string
	Volume float64
}

type MuscleVolumeDistribution struct {
	Muscle string
	Volume float64
}
