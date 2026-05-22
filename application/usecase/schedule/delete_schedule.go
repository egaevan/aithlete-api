package schedule

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type DeleteScheduleUseCase interface {
	Delete(ctx context.Context, userID, scheduleID string) error
}

type deleteScheduleUseCase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewDeleteScheduleUseCase(scheduleRepo repository.ScheduleRepository) DeleteScheduleUseCase {
	return &deleteScheduleUseCase{scheduleRepo: scheduleRepo}
}

func (uc *deleteScheduleUseCase) Delete(ctx context.Context, userID, scheduleID string) error {
	s, err := uc.scheduleRepo.FindByID(ctx, scheduleID)
	if err != nil {
		if service.IsScheduleNotFound(err) {
			return domainerr.ErrScheduleNotFound
		}
		return fmt.Errorf("find schedule: %w", err)
	}

	if s.UserID != userID {
		return domainerr.ErrScheduleNotFound
	}

	if err := uc.scheduleRepo.Delete(ctx, scheduleID); err != nil {
		return fmt.Errorf("delete schedule: %w", err)
	}

	return nil
}
