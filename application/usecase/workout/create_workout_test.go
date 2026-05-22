package workout

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestCreateWorkout_Success(t *testing.T) {
	workoutRepo := repository.NewMockWorkoutRepository()
	uc := NewCreateWorkoutUseCase(workoutRepo)
	ctx := context.Background()

	result, err := uc.Create(ctx, "user-1", "Upper Body", "2026-05-19", "lbs", "Felt strong")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "Upper Body" {
		t.Errorf("expected Name 'Upper Body', got %s", result.Name)
	}
	if result.UserID != "user-1" {
		t.Errorf("expected UserID 'user-1', got %s", result.UserID)
	}
	if result.Date != "2026-05-19" {
		t.Errorf("expected Date '2026-05-19', got %s", result.Date)
	}
	if result.WeightUnit != "lbs" {
		t.Errorf("expected WeightUnit 'lbs', got %s", result.WeightUnit)
	}
	if result.Notes != "Felt strong" {
		t.Errorf("expected Notes 'Felt strong', got %s", result.Notes)
	}
	if result.Completed {
		t.Error("expected new workout to be not completed")
	}
	if result.ID == "" {
		t.Error("expected ID to be set")
	}
	if result.CreatedAt == "" {
		t.Error("expected CreatedAt to be set")
	}
	if result.UpdatedAt == "" {
		t.Error("expected UpdatedAt to be set")
	}
}
