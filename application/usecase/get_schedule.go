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

type GetScheduleUseCase interface {
	Get(ctx context.Context, userID, scheduleID string) (*dto.ScheduleResult, error)
}

type getScheduleUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewGetScheduleUseCase(scheduleRepo repository.ScheduleRepository) GetScheduleUseCase {
	return &getScheduleUseCase{scheduleRepo: scheduleRepo}
}

func (uc *getScheduleUseCase) Get(ctx context.Context, userID, scheduleID string) (*dto.ScheduleResult, error) {
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

	return mapper.ScheduleToResult(s), nil
}
