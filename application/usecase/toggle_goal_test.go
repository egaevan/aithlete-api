package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestToggleGoal_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewToggleGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	result, err := uc.Toggle(ctx, "user-1", g.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.Completed {
		t.Error("expected goal to be completed after toggle")
	}
}

func TestToggleGoal_NotFound(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewToggleGoalUseCase(goalRepo)
	ctx := context.Background()

	_, err := uc.Toggle(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent goal")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}

func TestToggleGoal_WrongUser(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewToggleGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	_, err := uc.Toggle(ctx, "user-2", g.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}
