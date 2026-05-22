package entity

import (
	"testing"
	"time"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestNewWorkout(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "Felt strong")

	if w.UserID != "user-1" {
		t.Errorf("expected UserID user-1, got %s", w.UserID)
	}
	if w.Name != "Upper Body" {
		t.Errorf("expected Name Upper Body, got %s", w.Name)
	}
	if w.Completed {
		t.Error("expected new workout to be not completed")
	}
	if len(w.Exercises) != 0 {
		t.Error("expected new workout to have no exercises")
	}
	if w.CreatedAt.IsZero() || w.UpdatedAt.IsZero() {
		t.Error("expected timestamps to be set")
	}
}

func TestAddExercise(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	exercise := WorkoutExercise{
		ID: "we-1",
		Exercise: ExerciseRef{
			ID:   "ex-1",
			Name: "Bench Press",
		},
	}

	err := w.AddExercise(exercise)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(w.Exercises) != 1 {
		t.Errorf("expected 1 exercise, got %d", len(w.Exercises))
	}
}

func TestAddDuplicateExercise(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	exercise := WorkoutExercise{
		ID: "we-1",
		Exercise: ExerciseRef{
			ID:   "ex-1",
			Name: "Bench Press",
		},
	}

	w.AddExercise(exercise)
	err := w.AddExercise(exercise)
	if err != domainerr.ErrDuplicateExercise {
		t.Errorf("expected ErrDuplicateExercise, got %v", err)
	}
}

func TestTotalVolume(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []WorkoutExercise{
		{
			ID: "we-1",
			Sets: []Set{
				{ID: "s-1", Reps: 10, Weight: 135},
				{ID: "s-2", Reps: 8, Weight: 155},
			},
		},
		{
			ID: "we-2",
			Sets: []Set{
				{ID: "s-3", Reps: 8, Weight: 95},
			},
		},
	}

	expected := (10*135 + 8*155 + 8*95)
	volume := w.TotalVolume()
	if volume != float64(expected) {
		t.Errorf("expected volume %d, got %f", expected, volume)
	}
}

func TestCompleteWorkout(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []WorkoutExercise{
		{
			ID: "we-1",
			Sets: []Set{
				{ID: "s-1", Reps: 10, Weight: 135, Completed: true},
			},
		},
	}

	originalUpdatedAt := w.UpdatedAt
	time.Sleep(time.Nanosecond)

	err := w.Complete()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !w.Completed {
		t.Error("expected workout to be completed")
	}
	if !w.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after completion")
	}
}

func TestCompleteEmptyWorkout(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")

	err := w.Complete()
	if err != domainerr.ErrEmptyWorkout {
		t.Errorf("expected ErrEmptyWorkout, got %v", err)
	}
}

func TestUpdateSet(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []WorkoutExercise{
		{
			ID: "we-1",
			Exercise: ExerciseRef{ID: "ex-1"},
			Sets: []Set{
				{ID: "s-1", Reps: 10, Weight: 135, RPE: 7},
			},
		},
	}

	err := w.UpdateSet("ex-1", "s-1", 12, 140, 8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Exercises[0].Sets[0].Reps != 12 {
		t.Errorf("expected reps 12, got %d", w.Exercises[0].Sets[0].Reps)
	}
	if w.Exercises[0].Sets[0].Weight != 140 {
		t.Errorf("expected weight 140, got %f", w.Exercises[0].Sets[0].Weight)
	}
}

func TestUpdateSetInvalidValues(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w.Exercises = []WorkoutExercise{
		{
			ID: "we-1",
			Exercise: ExerciseRef{ID: "ex-1"},
			Sets: []Set{
				{ID: "s-1", Reps: 10, Weight: 135},
			},
		},
	}

	err := w.UpdateSet("ex-1", "s-1", 0, 140, 8)
	if err != domainerr.ErrInvalidSetValue {
		t.Errorf("expected ErrInvalidSetValue, got %v", err)
	}
}

func TestWorkoutUpdate(t *testing.T) {
	w := NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	originalUpdatedAt := w.UpdatedAt

	time.Sleep(time.Nanosecond)

	w.Update("Lower Body", "2026-05-20")
	if w.Name != "Lower Body" {
		t.Errorf("expected Name Lower Body, got %s", w.Name)
	}
	if w.Date != "2026-05-20" {
		t.Errorf("expected Date 2026-05-20, got %s", w.Date)
	}
	if !w.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after update")
	}
}
