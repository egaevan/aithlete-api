package entity

import (
	"testing"
	"time"
)

func TestNewBodyWeight(t *testing.T) {
	bw := NewBodyWeight("user-1", "2026-05-20", 185.5, 15.0)

	if bw.UserID != "user-1" {
		t.Errorf("expected UserID user-1, got %s", bw.UserID)
	}
	if bw.Weight != 185.5 {
		t.Errorf("expected Weight 185.5, got %f", bw.Weight)
	}
	if bw.BodyFatPercentage != 15.0 {
		t.Errorf("expected BodyFatPercentage 15.0, got %f", bw.BodyFatPercentage)
	}
	if bw.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}
}

func TestNewBodyWeightDefaults(t *testing.T) {
	bw := NewBodyWeight("user-1", "2026-05-20", 180.0, 0)

	if bw.BodyFatPercentage != 0 {
		t.Errorf("expected BodyFatPercentage 0, got %f", bw.BodyFatPercentage)
	}
	if !bw.CreatedAt.Before(time.Now()) && !bw.CreatedAt.Equal(time.Now()) {
		t.Error("expected CreatedAt to be recent")
	}
}
