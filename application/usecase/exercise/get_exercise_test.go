package exercise

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestGetExercise_Success(t *testing.T) {
	exerciseRepo := repository.NewMockExerciseRepository()
	uc := NewGetExerciseUseCase(exerciseRepo)
	ctx := context.Background()

	e := entity.NewExercise("Bench Press", "Barbell bench press", "chest", "barbell", "intermediate", nil)
	exerciseRepo.Create(ctx, e)

	result, err := uc.Get(ctx, e.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "Bench Press" {
		t.Errorf("expected name 'Bench Press', got %s", result.Name)
	}
	if result.MuscleGroup != "chest" {
		t.Errorf("expected muscleGroup 'chest', got %s", result.MuscleGroup)
	}
}

func TestGetExercise_NotFound(t *testing.T) {
	exerciseRepo := repository.NewMockExerciseRepository()
	uc := NewGetExerciseUseCase(exerciseRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent exercise")
	}
	if !errors.Is(err, domainerr.ErrExerciseNotFound) {
		t.Errorf("expected ErrExerciseNotFound, got %v", err)
	}
}
