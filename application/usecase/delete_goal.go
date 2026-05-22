package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type DeleteGoalUseCase interface {
	Delete(ctx context.Context, userID, goalID string) error
}

type deleteGoalUseCase struct {
	goalRepo repository.GoalRepository
}

func NewDeleteGoalUseCase(goalRepo repository.GoalRepository) DeleteGoalUseCase {
	return &deleteGoalUseCase{goalRepo: goalRepo}
}

func (uc *deleteGoalUseCase) Delete(ctx context.Context, userID, goalID string) error {
	g, err := uc.goalRepo.FindByID(ctx, goalID)
	if err != nil {
		if service.IsGoalNotFound(err) {
			return domainerr.ErrGoalNotFound
		}
		return fmt.Errorf("find goal: %w", err)
	}

	if g.UserID != userID {
		return domainerr.ErrGoalNotFound
	}

	if err := uc.goalRepo.Delete(ctx, goalID); err != nil {
		return fmt.Errorf("delete goal: %w", err)
	}

	return nil
}
