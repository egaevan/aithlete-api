package goal

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestDeleteGoal_Success(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewDeleteGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	err := uc.Delete(ctx, "user-1", g.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = goalRepo.FindByID(ctx, g.ID)
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Error("expected goal to be deleted")
	}
}

func TestDeleteGoal_NotFound(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewDeleteGoalUseCase(goalRepo)
	ctx := context.Background()

	err := uc.Delete(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent goal")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}

func TestDeleteGoal_WrongUser(t *testing.T) {
	goalRepo := repository.NewMockGoalRepository()
	uc := NewDeleteGoalUseCase(goalRepo)
	ctx := context.Background()

	g := entity.NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	goalRepo.Create(ctx, g)

	err := uc.Delete(ctx, "user-2", g.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrGoalNotFound) {
		t.Errorf("expected ErrGoalNotFound, got %v", err)
	}
}
