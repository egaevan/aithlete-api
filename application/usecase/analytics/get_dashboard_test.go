package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetDashboard_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetDashboardUseCase(analyticsRepo)
	ctx := context.Background()

	result, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Stats.CaloriesBurned != 2450 {
		t.Errorf("expected CaloriesBurned 2450, got %d", result.Stats.CaloriesBurned)
	}
	if len(result.WeeklyProgress) != 7 {
		t.Errorf("expected 7 weekly progress days, got %d", len(result.WeeklyProgress))
	}
	if result.Streak.Current != 5 {
		t.Errorf("expected Current streak 5, got %d", result.Streak.Current)
	}
}
