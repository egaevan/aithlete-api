package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateGoal_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	result, err := uc.Update(ctx, "user-1", g.ID, "Run 200 miles", "custom", 200, 50, "miles", "monthly", "2026-06-30")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Title != "Run 200 miles" {
		t.Errorf("expected Title 'Run 200 miles', got %s", result.Title)
	}
	if result.Target != 200 {
		t.Errorf("expected Target 200, got %d", result.Target)
	}
	if result.Current != 50 {
		t.Errorf("expected Current 50, got %d", result.Current)
	}
}

func TestUpdateGoal_NotFound(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalUseCase(goalRepo)
	ctx := context.Background()

	_, err := uc.Update(ctx, "user-1", "nonexistent", "Run 200 miles", "custom", 200, 50, "miles", "monthly", "2026-06-30")
	if err == nil {
		t.Fatal("expected error for nonexistent goal")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}

func TestUpdateGoal_WrongUser(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewUpdateGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	_, err := uc.Update(ctx, "user-2", g.ID, "Run 200 miles", "custom", 200, 50, "miles", "monthly", "2026-06-30")
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}
