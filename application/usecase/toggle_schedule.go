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

type ToggleScheduleUseCase interface {
	Toggle(ctx context.Context, userID, scheduleID string) (*dto.ScheduleResult, error)
}

type toggleScheduleUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewToggleScheduleUseCase(scheduleRepo repository.ScheduleRepository) ToggleScheduleUseCase {
	return &toggleScheduleUseCase{scheduleRepo: scheduleRepo}
}

func (uc *toggleScheduleUseCase) Toggle(ctx context.Context, userID, scheduleID string) (*dto.ScheduleResult, error) {
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

	s.Toggle()

	if err := uc.scheduleRepo.Update(ctx, s); err != nil {
		return nil, fmt.Errorf("update schedule: %w", err)
	}

	return mapper.ScheduleToResult(s), nil
}
