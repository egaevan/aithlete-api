package workout

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateWorkout_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	result, err := uc.Update(ctx, "user-1", w.ID, "Lower Body", "2026-05-20")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "Lower Body" {
		t.Errorf("expected Name 'Lower Body', got %s", result.Name)
	}
	if result.Date != "2026-05-20" {
		t.Errorf("expected Date '2026-05-20', got %s", result.Date)
	}
}

func TestUpdateWorkout_NotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	_, err := uc.Update(ctx, "user-1", "nonexistent", "Lower Body", "2026-05-20")
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}

func TestUpdateWorkout_WrongUser(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	_, err := uc.Update(ctx, "user-2", w.ID, "Lower Body", "2026-05-20")
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}
