package progress

import (
	"context"
	"testing"
	"time"

	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestAddBodyWeight_Success(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewAddBodyWeightUseCase(progressRepo)
	ctx := context.Background()

	result, err := uc.Add(ctx, "user-1", "2026-05-20", 185.5, 15.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Weight != 185.5 {
		t.Errorf("expected Weight 185.5, got %f", result.Weight)
	}
	if result.BodyFatPercentage != 15.0 {
		t.Errorf("expected BodyFatPercentage 15.0, got %f", result.BodyFatPercentage)
	}
	if result.UserID != "user-1" {
		t.Errorf("expected UserID 'user-1', got %s", result.UserID)
	}
	if result.ID == "" {
		t.Error("expected ID to be set")
	}
	if result.CreatedAt == "" {
		t.Error("expected CreatedAt to be set")
	}
}

func TestAddBodyWeight_DefaultDate(t *testing.T) {
	progressRepo := repository.NewMockProgressRepository()
	uc := NewAddBodyWeightUseCase(progressRepo)
	ctx := context.Background()

	result, err := uc.Add(ctx, "user-1", "", 180.0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := time.Now().Format("2006-01-02")
	if result.Date != expected {
		t.Errorf("expected Date %s, got %s", expected, result.Date)
	}
}
