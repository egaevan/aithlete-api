package dto

type BodyWeightResult struct {
	ID                string
	UserID            string
	Date              string
	Weight            float64
	BodyFatPercentage float64
	CreatedAt         string
}

type StrengthResult struct {
	ID        string
	UserID    string
	Exercise  string
	Date      string
	OneRepMax float64
	Volume    float64
}

type ConsistencyResult struct {
	Week              string
	WorkoutsCompleted int
	WorkoutsPlanned   int
	Streak            int
}

type MuscleVolumeResult struct {
	MuscleGroup string
	Volume      float64
	Sessions    int
	Trend       string
}

type ProgressOverviewResult struct {
	TotalWorkouts    int
	CurrentStreak    int
	LongestStreak    int
	TotalVolume      float64
	WorkoutsThisWeek int
	BodyWeight       *BodyWeightResult
}
