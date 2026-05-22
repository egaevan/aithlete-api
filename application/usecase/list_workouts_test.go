package usecase

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestListWorkouts_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewListWorkoutsUseCase(workoutRepo)
	ctx := context.Background()

	w1 := entity.NewWorkout("user-1", "Upper Body", "2026-05-19", "lbs", "")
	w2 := entity.NewWorkout("user-1", "Lower Body", "2026-05-20", "kg", "")
	w3 := entity.NewWorkout("user-2", "Cardio", "2026-05-21", "lbs", "")
	workoutRepo.Create(ctx, w1)
	workoutRepo.Create(ctx, w2)
	workoutRepo.Create(ctx, w3)

	results, err := uc.List(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 workouts, got %d", len(results))
	}

	if results[0].Name == "Cardio" && results[1].Name == "Cardio" {
		t.Error("expected only user-1 workouts")
	}
}

func TestListWorkouts_Empty(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewListWorkoutsUseCase(workoutRepo)
	ctx := context.Background()

	results, err := uc.List(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected empty list, got %d workouts", len(results))
	}
}
