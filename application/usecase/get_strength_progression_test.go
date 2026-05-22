package usecase

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetStrengthProgression_Success(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetStrengthProgressionUseCase(progressRepo)
	ctx := context.Background()

	progressRepo.StrengthRecords["sr-1"] = &entity.StrengthRecord{
		ID: "sr-1", UserID: "user-1", Exercise: "Bench Press", Date: "2026-05-01", OneRepMax: 225, Volume: 3000,
	}
	progressRepo.StrengthRecords["sr-2"] = &entity.StrengthRecord{
		ID: "sr-2", UserID: "user-1", Exercise: "Squat", Date: "2026-05-01", OneRepMax: 315, Volume: 5000,
	}

	results, err := uc.GetStrengthProgression(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 records, got %d", len(results))
	}
}

func TestGetStrengthProgression_Empty(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetStrengthProgressionUseCase(progressRepo)
	ctx := context.Background()

	results, err := uc.GetStrengthProgression(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected empty list, got %d records", len(results))
	}
}
