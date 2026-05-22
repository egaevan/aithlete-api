package analytics

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type GetStreakUseCase interface {
	Get(ctx context.Context, userID string) (*dto.StreakResult, error)
}

type getStreakUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetStreakUseCase(analyticsRepo repository.AnalyticsRepository) GetStreakUseCase {
	return &getStreakUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getStreakUseCase) Get(ctx context.Context, userID string) (*dto.StreakResult, error) {
	s, err := uc.analyticsRepo.GetStreak(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get streak: %w", err)
	}
	result := mapper.StreakToResult(s)
	return &result, nil
}
