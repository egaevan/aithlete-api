package workout

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestAddExercise_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewAddExerciseUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	sets := []entity.Set{
		{ID: "s-1", Reps: 10, Weight: 135, RPE: 7},
		{ID: "s-2", Reps: 8, Weight: 155, RPE: 8},
	}

	result, err := uc.AddExercise(ctx, "user-1", w.ID, entity.ExerciseRef{ID: "ex-1", Name: "Bench Press"}, sets)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Exercises) != 1 {
		t.Fatalf("expected 1 exercise, got %d", len(result.Exercises))
	}
	if result.Exercises[0].Exercise.Name != "Bench Press" {
		t.Errorf("expected exercise name 'Bench Press', got %s", result.Exercises[0].Exercise.Name)
	}
	if len(result.Exercises[0].Sets) != 2 {
		t.Errorf("expected 2 sets, got %d", len(result.Exercises[0].Sets))
	}
}

func TestAddExercise_Duplicate(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewAddExerciseUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	sets := []entity.Set{{ID: "s-1", Reps: 10, Weight: 135}}

	uc.AddExercise(ctx, "user-1", w.ID, entity.ExerciseRef{ID: "ex-1", Name: "Bench Press"}, sets)

	_, err := uc.AddExercise(ctx, "user-1", w.ID, entity.ExerciseRef{ID: "ex-1", Name: "Bench Press"}, sets)
	if err == nil {
		t.Fatal("expected error for duplicate exercise")
	}
	if !errors.Is(err, domainerr.ErrDuplicateExercise) {
		t.Errorf("expected ErrDuplicateExercise, got %v", err)
	}
}

func TestAddExercise_WorkoutNotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewAddExerciseUseCase(workoutRepo)
	ctx := context.Background()

	sets := []entity.Set{{ID: "s-1", Reps: 10, Weight: 135}}
	_, err := uc.AddExercise(ctx, "user-1", "nonexistent", entity.ExerciseRef{ID: "ex-1", Name: "Bench Press"}, sets)
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}
