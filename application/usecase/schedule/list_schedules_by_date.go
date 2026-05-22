package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type ListSchedulesByDateUseCase interface {
	ListByDate(ctx context.Context, userID, date string) ([]dto.ScheduleResult, error)
}

type listSchedulesByDateUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewListSchedulesByDateUseCase(scheduleRepo repository.ScheduleRepository) ListSchedulesByDateUseCase {
	return &listSchedulesByDateUseCase{scheduleRepo: scheduleRepo}
}

func (uc *listSchedulesByDateUseCase) ListByDate(ctx context.Context, userID, date string) ([]dto.ScheduleResult, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	schedules, err := uc.scheduleRepo.FindByUserIDAndDate(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("list schedules by date: %w", err)
	}

	results := make([]dto.ScheduleResult, len(schedules))
	for i := range schedules {
		results[i] = *mapper.ScheduleToResult(&schedules[i])
	}

	return results, nil
}
