package workout

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestGetWorkout_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewGetWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	result, err := uc.Get(ctx, "user-1", w.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ID != w.ID {
		t.Errorf("expected ID %s, got %s", w.ID, result.ID)
	}
	if result.UserID != "user-1" {
		t.Errorf("expected UserID 'user-1', got %s", result.UserID)
	}
}

func TestGetWorkout_NotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewGetWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}

func TestGetWorkout_WrongUser(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewGetWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	_, err := uc.Get(ctx, "user-2", w.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}
