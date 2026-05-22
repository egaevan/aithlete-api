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

type GetWeeklyVolumeUseCase interface {
	Get(ctx context.Context, userID string) ([]dto.WeeklyVolumeResult, error)
}

type getWeeklyVolumeUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetWeeklyVolumeUseCase(analyticsRepo repository.AnalyticsRepository) GetWeeklyVolumeUseCase {
	return &getWeeklyVolumeUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getWeeklyVolumeUseCase) Get(ctx context.Context, userID string) ([]dto.WeeklyVolumeResult, error) {
	results, err := uc.analyticsRepo.GetWeeklyVolume(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get weekly volume: %w", err)
	}
	return mapper.WeeklyVolumeToResultList(results), nil
}
