package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestToggleSchedule_Success(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewToggleScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	result, err := uc.Toggle(ctx, "user-1", s.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.Completed {
		t.Error("expected schedule to be completed after toggle")
	}
}

func TestToggleSchedule_NotFound(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewToggleScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	_, err := uc.Toggle(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent schedule")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}

func TestToggleSchedule_WrongUser(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewToggleScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	_, err := uc.Toggle(ctx, "user-2", s.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}
