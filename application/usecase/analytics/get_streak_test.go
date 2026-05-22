package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetStreak_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetStreakUseCase(analyticsRepo)
	ctx := context.Background()

	result, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Current != 5 {
		t.Errorf("expected Current streak 5, got %d", result.Current)
	}
	if result.PersonalBest != 14 {
		t.Errorf("expected PersonalBest 14, got %d", result.PersonalBest)
	}
}

func TestGetStreak_NoData(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetStreakUseCase(analyticsRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty userID")
	}
}
