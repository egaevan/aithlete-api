package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetWeeklyProgress_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetWeeklyProgressUseCase(analyticsRepo)
	ctx := context.Background()

	results, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected non-empty weekly progress")
	}
	if results[0].Day != "Mon" {
		t.Errorf("expected first day Mon, got %s", results[0].Day)
	}
}

func TestGetWeeklyProgress_NoData(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetWeeklyProgressUseCase(analyticsRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty userID")
	}
}
