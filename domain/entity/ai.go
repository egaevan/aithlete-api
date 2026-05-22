package entity

import "time"

type Recommendation struct {
	ID          string
	UserID      string
	Type        string
	Title       string
	Description string
	Confidence  float64
	CreatedAt   time.Time
}

type ChatSession struct {
	SessionID string
	UserID    string
	CreatedAt time.Time
}

type ChatMessage struct {
	ID        string
	Role      string
	Content   string
	Timestamp time.Time
}

type FatigueAnalysis struct {
	Overall    int
	Central    int
	Peripheral int
	Status     string
	Factors    []FatigueFactor
}

type FatigueFactor struct {
	Name   string
	Value  int
	Impact string
}

type RecoveryScore struct {
	Overall        int
	Sleep          int
	Nutrition      int
	Stress         int
	MuscleRecovery []MuscleRecovery
	Status         string
}

type PlateauDetection struct {
	Detected     bool
	Exercise     string
	Metric       string
	CurrentTrend string
	WeeksStalled int
	Suggestions  []string
}

func NewRecommendation(userID, rType, title, description string, confidence float64) *Recommendation {
	return &Recommendation{
		UserID:      userID,
		Type:        rType,
		Title:       title,
		Description: description,
		Confidence:  confidence,
		CreatedAt:   time.Now(),
	}
}

func NewChatSession(userID string) *ChatSession {
	return &ChatSession{
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}
