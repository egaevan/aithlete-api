package goal

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateGoalProgress_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalProgressUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	result, err := uc.UpdateProgress(ctx, "user-1", g.ID, 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Current != 50 {
		t.Errorf("expected Current 50, got %d", result.Current)
	}
	if result.Completed {
		t.Error("expected goal to not be completed yet")
	}
}

func TestUpdateGoalProgress_CompletesGoal(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalProgressUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	result, err := uc.UpdateProgress(ctx, "user-1", g.ID, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.Completed {
		t.Error("expected goal to be completed when current reaches target")
	}
}

func TestUpdateGoalProgress_NotFound(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalProgressUseCase(goalRepo)
	ctx := context.Background()

	_, err := uc.UpdateProgress(ctx, "user-1", "nonexistent", 50)
	if err == nil {
		t.Fatal("expected error for nonexistent goal")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}

func TestUpdateGoalProgress_WrongUser(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalProgressUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	_, err := uc.UpdateProgress(ctx, "user-2", g.ID, 50)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}
