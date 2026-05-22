package progress

import "github.com/aithlete/aithlete-api/application/dto"

type BodyWeightResponse struct {
	ID                string  `json:"id"`
	UserID            string  `json:"userId"`
	Date              string  `json:"date"`
	Weight            float64 `json:"weight"`
	BodyFatPercentage float64 `json:"bodyFatPercentage"`
	CreatedAt         string  `json:"createdAt"`
}

type StrengthResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userId"`
	Exercise  string  `json:"exercise"`
	Date      string  `json:"date"`
	OneRepMax float64 `json:"oneRepMax"`
	Volume    float64 `json:"volume"`
}

type ConsistencyResponse struct {
	Week              string `json:"week"`
	WorkoutsCompleted int    `json:"workoutsCompleted"`
	WorkoutsPlanned   int    `json:"workoutsPlanned"`
	Streak            int    `json:"streak"`
}

type MuscleVolumeResponse struct {
	MuscleGroup string  `json:"muscleGroup"`
	Volume      float64 `json:"volume"`
	Sessions    int     `json:"sessions"`
	Trend       string  `json:"trend"`
}

type ProgressOverviewResponse struct {
	TotalWorkouts    int                 `json:"totalWorkouts"`
	CurrentStreak    int                 `json:"currentStreak"`
	LongestStreak    int                 `json:"longestStreak"`
	TotalVolume      float64             `json:"totalVolume"`
	WorkoutsThisWeek int                 `json:"workoutsThisWeek"`
	BodyWeight       *BodyWeightResponse `json:"bodyWeight,omitempty"`
}

func toBodyWeightResponse(r *dto.BodyWeightResult) BodyWeightResponse {
	return BodyWeightResponse{
		ID: r.ID, UserID: r.UserID, Date: r.Date,
		Weight: r.Weight, BodyFatPercentage: r.BodyFatPercentage, CreatedAt: r.CreatedAt,
	}
}

func toBodyWeightResponseList(results []dto.BodyWeightResult) []BodyWeightResponse {
	resp := make([]BodyWeightResponse, len(results))
	for i, r := range results {
		resp[i] = toBodyWeightResponse(&r)
	}
	return resp
}

func toStrengthResponseList(results []dto.StrengthResult) []StrengthResponse {
	resp := make([]StrengthResponse, len(results))
	for i, r := range results {
		resp[i] = StrengthResponse{
			ID: r.ID, UserID: r.UserID, Exercise: r.Exercise, Date: r.Date,
			OneRepMax: r.OneRepMax, Volume: r.Volume,
		}
	}
	return resp
}

func toConsistencyResponseList(results []dto.ConsistencyResult) []ConsistencyResponse {
	resp := make([]ConsistencyResponse, len(results))
	for i, r := range results {
		resp[i] = ConsistencyResponse{
			Week: r.Week, WorkoutsCompleted: r.WorkoutsCompleted,
			WorkoutsPlanned: r.WorkoutsPlanned, Streak: r.Streak,
		}
	}
	return resp
}

func toMuscleVolumeResponseList(results []dto.MuscleVolumeResult) []MuscleVolumeResponse {
	resp := make([]MuscleVolumeResponse, len(results))
	for i, r := range results {
		resp[i] = MuscleVolumeResponse{
			MuscleGroup: r.MuscleGroup, Volume: r.Volume,
			Sessions: r.Sessions, Trend: r.Trend,
		}
	}
	return resp
}

func toProgressOverviewResponse(r *dto.ProgressOverviewResult) ProgressOverviewResponse {
	resp := ProgressOverviewResponse{
		TotalWorkouts: r.TotalWorkouts, CurrentStreak: r.CurrentStreak,
		LongestStreak: r.LongestStreak, TotalVolume: r.TotalVolume,
		WorkoutsThisWeek: r.WorkoutsThisWeek,
	}
	if r.BodyWeight != nil {
		bw := toBodyWeightResponse(r.BodyWeight)
		resp.BodyWeight = &bw
	}
	return resp
}
