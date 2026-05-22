package progress

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetProgressOverview_Success(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetProgressOverviewUseCase(progressRepo)
	ctx := context.Background()

	progressRepo.AddBodyWeight(ctx, entity.NewBodyWeight("user-1", "2026-05-20", 185.5, 15.0))
	progressRepo.StrengthRecords["sr-1"] = &entity.StrengthRecord{
		ID: "sr-1", UserID: "user-1", Exercise: "Bench Press", Volume: 3000,
	}

	result, err := uc.GetOverview(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalVolume != 3000 {
		t.Errorf("expected TotalVolume 3000, got %f", result.TotalVolume)
	}
	if result.BodyWeight == nil {
		t.Fatal("expected BodyWeight to be set")
	}
	if result.BodyWeight.Weight != 185.5 {
		t.Errorf("expected BodyWeight.Weight 185.5, got %f", result.BodyWeight.Weight)
	}
}

func TestGetProgressOverview_Empty(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetProgressOverviewUseCase(progressRepo)
	ctx := context.Background()

	result, err := uc.GetOverview(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.BodyWeight != nil {
		t.Error("expected nil BodyWeight for empty data")
	}
}
