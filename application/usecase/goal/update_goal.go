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

type UpdateGoalUseCase interface {
	Update(ctx context.Context, userID, goalID, title, typ string, target, current int, unit, period, deadline string) (*dto.GoalResult, error)
}

type updateGoalUseCase struct {
	goalRepo repository.GoalRepository
}

func NewUpdateGoalUseCase(goalRepo repository.GoalRepository) UpdateGoalUseCase {
	return &updateGoalUseCase{goalRepo: goalRepo}
}

func (uc *updateGoalUseCase) Update(ctx context.Context, userID, goalID, title, typ string, target, current int, unit, period, deadline string) (*dto.GoalResult, error) {
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

	g.Update(title, typ, target, current, unit, period, deadline)

	if err := uc.goalRepo.Update(ctx, g); err != nil {
		return nil, fmt.Errorf("update goal: %w", err)
	}

	return mapper.GoalToResult(g), nil
}
