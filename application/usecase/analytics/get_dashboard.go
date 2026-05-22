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

type GetDashboardUseCase interface {
	Get(ctx context.Context, userID string) (*dto.DashboardResult, error)
}

type getDashboardUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetDashboardUseCase(analyticsRepo repository.AnalyticsRepository) GetDashboardUseCase {
	return &getDashboardUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getDashboardUseCase) Get(ctx context.Context, userID string) (*dto.DashboardResult, error) {
	d, err := uc.analyticsRepo.GetDashboard(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get dashboard: %w", err)
	}
	return mapper.DashboardToResult(d), nil
}
