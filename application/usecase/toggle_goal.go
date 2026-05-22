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

type ToggleGoalUseCase interface {
	Toggle(ctx context.Context, userID, goalID string) (*dto.GoalResult, error)
}

type toggleGoalUseCase struct {
	goalRepo repository.GoalRepository
}

func NewToggleGoalUseCase(goalRepo repository.GoalRepository) ToggleGoalUseCase {
	return &toggleGoalUseCase{goalRepo: goalRepo}
}

func (uc *toggleGoalUseCase) Toggle(ctx context.Context, userID, goalID string) (*dto.GoalResult, error) {
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

	g.Toggle()

	if err := uc.goalRepo.Update(ctx, g); err != nil {
		return nil, fmt.Errorf("update goal: %w", err)
	}

	return mapper.GoalToResult(g), nil
}
