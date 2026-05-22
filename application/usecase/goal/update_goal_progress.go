package goal

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type UpdateGoalProgressUseCase interface {
	UpdateProgress(ctx context.Context, userID, goalID string, current int) (*dto.GoalResult, error)
}

type updateGoalProgressUseCase struct {
	goalRepo repository.GoalRepository
}

func NewUpdateGoalProgressUseCase(goalRepo repository.GoalRepository) UpdateGoalProgressUseCase {
	return &updateGoalProgressUseCase{goalRepo: goalRepo}
}

func (uc *updateGoalProgressUseCase) UpdateProgress(ctx context.Context, userID, goalID string, current int) (*dto.GoalResult, error) {
	g, err := uc.goalRepo.FindByID(ctx, goalID)
	if err != nil {
		if service.IsGoalNotFound(err) {
			return nil, domainerr.ErrGoalNotFound
		}
		return nil, fmt.Errorf("find goal: %w", err)
	}

	if g.UserID != userID {
		return nil, domainerr.ErrGoalNotFound
	}

	g.UpdateProgress(current)

	if err := uc.goalRepo.Update(ctx, g); err != nil {
		return nil, fmt.Errorf("update goal: %w", err)
	}

	return mapper.GoalToResult(g), nil
}
