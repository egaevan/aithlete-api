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

type GetWeeklyProgressUseCase interface {
	Get(ctx context.Context, userID string) ([]dto.WeeklyProgressDayResult, error)
}

type getWeeklyProgressUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetWeeklyProgressUseCase(analyticsRepo repository.AnalyticsRepository) GetWeeklyProgressUseCase {
	return &getWeeklyProgressUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getWeeklyProgressUseCase) Get(ctx context.Context, userID string) ([]dto.WeeklyProgressDayResult, error) {
	results, err := uc.analyticsRepo.GetWeeklyProgress(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get weekly progress: %w", err)
	}
	return mapper.WeeklyProgressDayToResultList(results), nil
}
