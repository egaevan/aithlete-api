package schedule

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type ListSchedulesUseCase interface {
	List(ctx context.Context, userID string) ([]dto.ScheduleResult, error)
}

type listSchedulesUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewListSchedulesUseCase(scheduleRepo repository.ScheduleRepository) ListSchedulesUseCase {
	return &listSchedulesUseCase{scheduleRepo: scheduleRepo}
}

func (uc *listSchedulesUseCase) List(ctx context.Context, userID string) ([]dto.ScheduleResult, error) {
	schedules, err := uc.scheduleRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list schedules: %w", err)
	}

	results := make([]dto.ScheduleResult, len(schedules))
	for i := range schedules {
		results[i] = *mapper.ScheduleToResult(&schedules[i])
	}

	return results, nil
}
