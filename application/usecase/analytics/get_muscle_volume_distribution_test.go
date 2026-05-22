package analytics

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetMuscleVolumeDistribution_Success(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetMuscleVolumeDistributionUseCase(analyticsRepo)
	ctx := context.Background()

	results, err := uc.Get(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected non-empty muscle volume distribution")
	}
	if results[0].Muscle != "chest" {
		t.Errorf("expected first muscle chest, got %s", results[0].Muscle)
	}
}

func TestGetMuscleVolumeDistribution_NoData(t *testing.T) {
	analyticsRepo := repository.NewMockAnalyticsRepository()
	uc := NewGetMuscleVolumeDistributionUseCase(analyticsRepo)
	ctx := context.Background()

	_, err := uc.Get(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty userID")
	}
}
