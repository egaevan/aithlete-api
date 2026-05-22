package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetOverview_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetOverviewUseCase(analyticsRepo)
	ctx := context.Background()

	result, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalVolume != 80100 {
		t.Errorf("expected TotalVolume 80100, got %d", result.TotalVolume)
	}
	if result.GoalCompletion != 20 {
		t.Errorf("expected GoalCompletion 20, got %d", result.GoalCompletion)
	}
}

func TestGetOverview_NoData(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetOverviewUseCase(analyticsRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty userID")
	}
}
