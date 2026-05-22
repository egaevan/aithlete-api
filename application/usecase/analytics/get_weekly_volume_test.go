package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetWeeklyVolume_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetWeeklyVolumeUseCase(analyticsRepo)
	ctx := context.Background()

	results, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected non-empty weekly volume")
	}
}

func TestGetWeeklyVolume_NoData(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetWeeklyVolumeUseCase(analyticsRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty userID")
	}
}
