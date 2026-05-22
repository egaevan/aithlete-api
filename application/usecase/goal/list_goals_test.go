package goal

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestListGoals_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewListGoalsUseCase(goalRepo)
	ctx := context.Background()

	goalRepo.Create(ctx, entity.NewGoal("user-1", "Goal 1", "custom", 10, "miles", "monthly", "2026-05-31"))
	goalRepo.Create(ctx, entity.NewGoal("user-1", "Goal 2", "custom", 20, "miles", "monthly", "2026-06-30"))
	goalRepo.Create(ctx, entity.NewGoal("user-2", "Other Goal", "custom", 30, "miles", "monthly", "2026-05-31"))

	results, err := uc.List(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 goals, got %d", len(results))
	}

	if results[0].Title != "Goal 1" && results[1].Title != "Goal 1" {
		t.Error("expected Goal 1 in results")
	}
	if results[0].Title != "Goal 2" && results[1].Title != "Goal 2" {
		t.Error("expected Goal 2 in results")
	}
}

func TestListGoals_Empty(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewListGoalsUseCase(goalRepo)
	ctx := context.Background()

	results, err := uc.List(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected 0 goals, got %d", len(results))
	}
}
