package progress

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestGetMuscleVolume_Empty(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewGetMuscleVolumeUseCase(progressRepo)
	ctx := context.Background()

	results, err := uc.GetMuscleVolume(ctx, "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected empty list, got %d records", len(results))
	}
}
