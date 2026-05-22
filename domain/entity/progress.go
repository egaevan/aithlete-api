package entity

import "time"

type BodyWeight struct {
	ID                string
	UserID            string
	Date              string
	Weight            float64
	BodyFatPercentage float64
	CreatedAt         time.Time
}

type StrengthRecord struct {
	ID        string
	UserID    string
	Exercise  string
	Date      string
	OneRepMax float64
	Volume    float64
}

type Consistency struct {
	Week              string
	WorkoutsCompleted int
	WorkoutsPlanned   int
	Streak            int
}

type MuscleVolume struct {
	MuscleGroup string
	Volume      float64
	Sessions    int
	Trend       string
}

func NewBodyWeight(userID, date string, weight, bodyFat float64) *BodyWeight {
	return &BodyWeight{
		UserID:            userID,
		Date:              date,
		Weight:            weight,
		BodyFatPercentage: bodyFat,
		CreatedAt:         time.Now(),
	}
}
