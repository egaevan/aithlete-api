package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestCompleteWorkout_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewCompleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []entity.WorkoutExercise{
		{
			ID: "we-1",
			Exercise: entity.ExerciseRef{ID: "ex-1"},
			Sets: []entity.Set{
				{ID: "s-1", Reps: 10, Weight: 135, Completed: true},
			},
		},
	}
	workoutRepo.Create(ctx, w)

	result, err := uc.Complete(ctx, "user-1", w.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.Completed {
		t.Error("expected workout to be completed")
	}
}

func TestCompleteWorkout_EmptyWorkout(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewCompleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	_, err := uc.Complete(ctx, "user-1", w.ID)
	if err == nil {
		t.Fatal("expected error for empty workout")
	}
	if !errors.Is(err, domainerr.ErrEmptyWorkout) {
		t.Errorf("expected ErrEmptyWorkout, got %v", err)
	}
}

func TestCompleteWorkout_NotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewCompleteWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	_, err := uc.Complete(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}
