package usecase

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetBodyWeightHistory_Success(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetBodyWeightHistoryUseCase(progressRepo)
	ctx := context.Background()

	bw1 := entity.NewBodyWeight("user-1", "2026-05-20", 185.5, 15.0)
	bw2 := entity.NewBodyWeight("user-1", "2026-05-27", 184.0, 14.5)
	progressRepo.AddBodyWeight(ctx, bw1)
	progressRepo.AddBodyWeight(ctx, bw2)

	results, err := uc.GetBodyWeightHistory(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 records, got %d", len(results))
	}
}

func TestGetBodyWeightHistory_Empty(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetBodyWeightHistoryUseCase(progressRepo)
	ctx := context.Background()

	results, err := uc.GetBodyWeightHistory(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected empty list, got %d records", len(results))
	}
}
