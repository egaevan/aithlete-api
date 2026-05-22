package schedule

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type CreateScheduleUseCase interface {
	Create(ctx context.Context, userID, date, schedTime, title, duration, typ, notes string) (*dto.ScheduleResult, error)
}

type createScheduleUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewCreateScheduleUseCase(scheduleRepo repository.ScheduleRepository) CreateScheduleUseCase {
	return &createScheduleUseCase{scheduleRepo: scheduleRepo}
}

func (uc *createScheduleUseCase) Create(ctx context.Context, userID, date, schedTime, title, duration, typ, notes string) (*dto.ScheduleResult, error) {
	s := entity.NewSchedule(userID, date, schedTime, title, duration, typ, notes)
	if err := uc.scheduleRepo.Create(ctx, s); err != nil {
		return nil, fmt.Errorf("create schedule: %w", err)
	}
	return mapper.ScheduleToResult(s), nil
}
