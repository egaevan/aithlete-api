package workout

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestDeleteWorkout_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewDeleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	err := uc.Delete(ctx, "user-1", w.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = workoutRepo.FindByID(ctx, w.ID)
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected workout to be deleted, got %v", err)
	}
}

func TestDeleteWorkout_NotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewDeleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	err := uc.Delete(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}

func TestDeleteWorkout_WrongUser(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewDeleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	err := uc.Delete(ctx, "user-2", w.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}
