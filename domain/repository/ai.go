package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type AIRepository interface {
	GetRecommendations(ctx context.Context, userID string) ([]entity.Recommendation, error)
	CreateChatSession(ctx context.Context, session *entity.ChatSession) error
	GetChatHistory(ctx context.Context, sessionID string) ([]entity.ChatMessage, error)
	SendChatMessage(ctx context.Context, sessionID string, message *entity.ChatMessage) ([]entity.ChatMessage, error)
	GetFatigueAnalysis(ctx context.Context, userID string) (*entity.FatigueAnalysis, error)
	GetRecoveryScore(ctx context.Context, userID string) (*entity.RecoveryScore, error)
	GetPlateauDetection(ctx context.Context, userID string) ([]entity.PlateauDetection, error)
}
