package goal

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestCreateGoal_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewCreateGoalUseCase(goalRepo)
	ctx := context.Background()

	result, err := uc.Create(ctx, "user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Title != "Run 100 miles" {
		t.Errorf("expected Title 'Run 100 miles', got %s", result.Title)
	}
	if result.UserID != "user-1" {
		t.Errorf("expected UserID 'user-1', got %s", result.UserID)
	}
	if result.Target != 100 {
		t.Errorf("expected Target 100, got %d", result.Target)
	}
	if result.Current != 0 {
		t.Errorf("expected Current 0, got %d", result.Current)
	}
	if result.Completed {
		t.Error("expected new goal to be not completed")
	}
	if result.ID == "" {
		t.Error("expected ID to be set")
	}
}
