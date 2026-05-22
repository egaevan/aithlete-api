package goal

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type ListGoalsUseCase interface {
	List(ctx context.Context, userID string) ([]dto.GoalResult, error)
}

type listGoalsUseCase struct {
	goalRepo repository.GoalRepository
}

func NewListGoalsUseCase(goalRepo repository.GoalRepository) ListGoalsUseCase {
	return &listGoalsUseCase{goalRepo: goalRepo}
}

func (uc *listGoalsUseCase) List(ctx context.Context, userID string) ([]dto.GoalResult, error) {
	goals, err := uc.goalRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list goals: %w", err)
	}
	return mapper.GoalToResultList(goals), nil
}
