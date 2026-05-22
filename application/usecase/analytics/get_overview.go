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

type GetOverviewUseCase interface {
	Get(ctx context.Context, userID string) (*dto.AnalyticsOverviewResult, error)
}

type getOverviewUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetOverviewUseCase(analyticsRepo repository.AnalyticsRepository) GetOverviewUseCase {
	return &getOverviewUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getOverviewUseCase) Get(ctx context.Context, userID string) (*dto.AnalyticsOverviewResult, error) {
	o, err := uc.analyticsRepo.GetOverview(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get overview: %w", err)
	}
	return mapper.AnalyticsOverviewToResult(o), nil
}
