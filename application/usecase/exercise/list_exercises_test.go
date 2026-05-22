package exercise

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestListExercises_Success(t *testing.T) {
	exerciseRepo := repository.NewMockExerciseRepository()
	uc := NewListExercisesUseCase(exerciseRepo)
	ctx := context.Background()

	exerciseRepo.Create(ctx, entity.NewExercise("Bench Press", "Barbell bench press", "chest", "barbell", "intermediate", nil))
	exerciseRepo.Create(ctx, entity.NewExercise("Squat", "Barbell squat", "legs", "barbell", "intermediate", nil))

	results, err := uc.List(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 exercises, got %d", len(results))
	}
}

func TestListExercises_Empty(t *testing.T) {
	exerciseRepo := repository.NewMockExerciseRepository()
	uc := NewListExercisesUseCase(exerciseRepo)
	ctx := context.Background()

	results, err := uc.List(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected 0 exercises, got %d", len(results))
	}
}
