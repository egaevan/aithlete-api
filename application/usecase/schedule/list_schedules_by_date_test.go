package schedule

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

func TestListSchedulesByDate_Success(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewListSchedulesByDateUseCase(scheduleRepo)
	ctx := context.Background()

	s1 := entity.NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	s2 := entity.NewSchedule("user-1", "2026-05-20", "07:00", "Yoga", "40 min", "stretching", "")
	s3 := entity.NewSchedule("user-1", "2026-05-21", "08:00", "Swimming", "45 min", "cardio", "")
	scheduleRepo.Create(ctx, s1)
	scheduleRepo.Create(ctx, s2)
	scheduleRepo.Create(ctx, s3)

	results, err := uc.ListByDate(ctx, "user-1", "2026-05-20")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 schedules, got %d", len(results))
	}
}

func TestListSchedulesByDate_Empty(t *testing.T) {
	scheduleRepo := repository.NewMockScheduleRepository()
	uc := NewListSchedulesByDateUseCase(scheduleRepo)
	ctx := context.Background()

	results, err := uc.ListByDate(ctx, "user-1", "2026-05-20")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected empty list, got %d schedules", len(results))
	}
}
