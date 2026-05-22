package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type UpdateScheduleUseCase interface {
	Update(ctx context.Context, userID, scheduleID, date, schedTime, title, duration, typ, notes string) (*dto.ScheduleResult, error)
}

type updateScheduleUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewUpdateScheduleUseCase(scheduleRepo repository.ScheduleRepository) UpdateScheduleUseCase {
	return &updateScheduleUseCase{scheduleRepo: scheduleRepo}
}

func (uc *updateScheduleUseCase) Update(ctx context.Context, userID, scheduleID, date, schedTime, title, duration, typ, notes string) (*dto.ScheduleResult, error) {
	s, err := uc.scheduleRepo.FindByID(ctx, scheduleID)
	if err != nil {
		if service.IsScheduleNotFound(err) {
			return nil, domainerr.ErrScheduleNotFound
		}
		return nil, fmt.Errorf("find schedule: %w", err)
	}

	if s.UserID != userID {
		return nil, domainerr.ErrScheduleNotFound
	}

	s.Update(date, schedTime, title, duration, typ, notes)

	if err := uc.scheduleRepo.Update(ctx, s); err != nil {
		return nil, fmt.Errorf("update schedule: %w", err)
	}

	return mapper.ScheduleToResult(s), nil
}
