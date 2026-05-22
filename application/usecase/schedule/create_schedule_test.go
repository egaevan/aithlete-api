package schedule

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestCreateSchedule_Success(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewCreateScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	result, err := uc.Create(ctx, "user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Title != "Morning Run" {
		t.Errorf("expected Title 'Morning Run', got %s", result.Title)
	}
	if result.UserID != "user-1" {
		t.Errorf("expected UserID 'user-1', got %s", result.UserID)
	}
	if result.Date != "2026-05-20" {
		t.Errorf("expected Date '2026-05-20', got %s", result.Date)
	}
	if result.Completed {
		t.Error("expected new schedule to be not completed")
	}
	if result.ID == "" {
		t.Error("expected ID to be set")
	}
}
