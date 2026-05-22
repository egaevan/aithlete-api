package schedule

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestDeleteSchedule_Success(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewDeleteScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	err := uc.Delete(ctx, "user-1", s.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = scheduleRepo.FindByID(ctx, s.ID)
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected schedule to be deleted, got %v", err)
	}
}

func TestDeleteSchedule_NotFound(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewDeleteScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	err := uc.Delete(ctx, "user-1", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent schedule")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}

func TestDeleteSchedule_WrongUser(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewDeleteScheduleUseCase(scheduleRepo)
	ctx := context.Background()

	s := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	scheduleRepo.Create(ctx, s)

	err := uc.Delete(ctx, "user-2", s.ID)
	if err == nil {
		t.Fatal("expected error for wrong user")
	}
	if !errors.Is(err, domainerr.ErrScheduleNotFound) {
		t.Errorf("expected ErrScheduleNotFound, got %v", err)
	}
}
