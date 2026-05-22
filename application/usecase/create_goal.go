package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type CreateGoalUseCase interface {
	Create(ctx context.Context, userID, title, typ string, target int, unit, period, deadline string) (*dto.GoalResult, error)
}

type createGoalUseCase struct {
	goalRepo repository.GoalRepository
}

func NewCreateGoalUseCase(goalRepo repository.GoalRepository) CreateGoalUseCase {
	return &createGoalUseCase{goalRepo: goalRepo}
}

func (uc *createGoalUseCase) Create(ctx context.Context, userID, title, typ string, target int, unit, period, deadline string) (*dto.GoalResult, error) {
	g := entity.NewGoal(userID, title, typ, target, unit, period, deadline)
	if err := uc.goalRepo.Create(ctx, g); err != nil {
		return nil, fmt.Errorf("create goal: %w", err)
	}
	return mapper.GoalToResult(g), nil
}
