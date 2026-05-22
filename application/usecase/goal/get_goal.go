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

type GetGoalUseCase interface {
	Get(ctx context.Context, userID, goalID string) (*dto.GoalResult, error)
}

type getGoalUseCase struct {
	goalRepo repository.GoalRepository
}

func NewGetGoalUseCase(goalRepo repository.GoalRepository) GetGoalUseCase {
	return &getGoalUseCase{goalRepo: goalRepo}
}

func (uc *getGoalUseCase) Get(ctx context.Context, userID, goalID string) (*dto.GoalResult, error) {
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

	return mapper.GoalToResult(g), nil
}
