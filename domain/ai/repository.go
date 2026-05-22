package ai

import "context"

type Repository interface {
	GetRecommendations(ctx context.Context, userID string) ([]Recommendation, error)
	CreateChatSession(ctx context.Context, session *ChatSession) error
	GetChatHistory(ctx context.Context, sessionID string) ([]ChatMessage, error)
	SendChatMessage(ctx context.Context, sessionID string, message *ChatMessage) ([]ChatMessage, error)
	GetFatigueAnalysis(ctx context.Context, userID string) (*FatigueAnalysis, error)
	GetRecoveryScore(ctx context.Context, userID string) (*RecoveryScore, error)
	GetPlateauDetection(ctx context.Context, userID string) ([]PlateauDetection, error)
}
