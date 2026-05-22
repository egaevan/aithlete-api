package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateSchedule_Success(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewUpdateScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	result, err := uc.Update(ctx, "user-1", s.ID, "2026-05-21", "07:00", "Yoga", "40 min", "stretching", "Relaxing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Title != "Yoga" {
		t.Errorf("expected Title 'Yoga', got %s", result.Title)
	}
	if result.Date != "2026-05-21" {
		t.Errorf("expected Date '2026-05-21', got %s", result.Date)
	}
	if result.Notes != "Relaxing" {
		t.Errorf("expected Notes 'Relaxing', got %s", result.Notes)
	}
}

func TestUpdateSchedule_NotFound(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewUpdateScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	_, err := uc.Update(ctx, "user-1", "nonexistent", "2026-05-21", "07:00", "Yoga", "40 min", "stretching", "")
	if err == nil {
		t.Fatal("expected error for nonexistent schedule")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}

func TestUpdateSchedule_WrongUser(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewUpdateScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	_, err := uc.Update(ctx, "user-2", s.ID, "2026-05-21", "07:00", "Yoga", "40 min", "stretching", "")
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}
