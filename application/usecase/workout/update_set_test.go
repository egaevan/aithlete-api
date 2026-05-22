package workout

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateSet_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateSetUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []entity.WorkoutExercise{
		{
			ID: "we-1",
			Exercise: entity.ExerciseRef{ID: "ex-1"},
			Sets: []entity.Set{
				{ID: "s-1", Reps: 10, Weight: 135, RPE: 7},
			},
		},
	}
	workoutRepo.Create(ctx, w)

	result, err := uc.UpdateSet(ctx, "user-1", w.ID, "ex-1", "s-1", 12, 140, 8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Exercises) != 1 {
		t.Fatalf("expected 1 exercise, got %d", len(result.Exercises))
	}
	if len(result.Exercises[0].Sets) != 1 {
		t.Fatalf("expected 1 set, got %d", len(result.Exercises[0].Sets))
	}
	if result.Exercises[0].Sets[0].Reps != 12 {
		t.Errorf("expected reps 12, got %d", result.Exercises[0].Sets[0].Reps)
	}
	if result.Exercises[0].Sets[0].Weight != 140 {
		t.Errorf("expected weight 140, got %f", result.Exercises[0].Sets[0].Weight)
	}
}

func TestUpdateSet_InvalidValues(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateSetUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []entity.WorkoutExercise{
		{
			ID: "we-1",
			Exercise: entity.ExerciseRef{ID: "ex-1"},
			Sets: []entity.Set{
				{ID: "s-1", Reps: 10, Weight: 135},
			},
		},
	}
	workoutRepo.Create(ctx, w)

	_, err := uc.UpdateSet(ctx, "user-1", w.ID, "ex-1", "s-1", 0, 140, 8)
	if err == nil {
		t.Fatal("expected error for invalid reps")
	}
	if !errors.Is(err, domainerr.ErrInvalidSetValue) {
		t.Errorf("expected ErrInvalidSetValue, got %v", err)
	}
}

func TestUpdateSet_WorkoutNotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateSetUseCase(workoutRepo)
	ctx := context.Background()

	_, err := uc.UpdateSet(ctx, "user-1", "nonexistent", "ex-1", "s-1", 12, 140, 8)
	if err == nil {
		t.Fatal("expected error for nonexistent workout")
	}
	if !errors.Is(err, domainerr.ErrWorkoutNotFound) {
		t.Errorf("expected ErrWorkoutNotFound, got %v", err)
	}
}

func TestUpdateSet_ExerciseNotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateSetUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	workoutRepo.Create(ctx, w)

	_, err := uc.UpdateSet(ctx, "user-1", w.ID, "nonexistent-exercise", "s-1", 12, 140, 8)
	if err == nil {
		t.Fatal("expected error for nonexistent exercise")
	}
	if !errors.Is(err, domainerr.ErrExerciseNotFoundInWorkout) {
		t.Errorf("expected ErrExerciseNotFoundInWorkout, got %v", err)
	}
}

func TestUpdateSet_SetNotFound(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewUpdateSetUseCase(workoutRepo)
	ctx := context.Background()

	w := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []entity.WorkoutExercise{
		{
			ID: "we-1",
			Exercise: entity.ExerciseRef{ID: "ex-1"},
			Sets: []entity.Set{
				{ID: "s-1", Reps: 10, Weight: 135},
			},
		},
	}
	workoutRepo.Create(ctx, w)

	_, err := uc.UpdateSet(ctx, "user-1", w.ID, "ex-1", "nonexistent-set", 12, 140, 8)
	if err == nil {
		t.Fatal("expected error for nonexistent set")
	}
	if !errors.Is(err, domainerr.ErrSetNotFound) {
		t.Errorf("expected ErrSetNotFound, got %v", err)
	}
}
