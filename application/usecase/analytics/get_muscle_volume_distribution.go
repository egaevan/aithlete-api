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

type GetMuscleVolumeDistributionUseCase interface {
	Get(ctx context.Context, userID string) ([]dto.MuscleVolumeDistributionResult, error)
}

type getMuscleVolumeDistributionUseCase struct {
	analyticsRepo repository.AnalyticsRepository
}

func NewGetMuscleVolumeDistributionUseCase(analyticsRepo repository.AnalyticsRepository) GetMuscleVolumeDistributionUseCase {
	return &getMuscleVolumeDistributionUseCase{analyticsRepo: analyticsRepo}
}

func (uc *getMuscleVolumeDistributionUseCase) Get(ctx context.Context, userID string) ([]dto.MuscleVolumeDistributionResult, error) {
	results, err := uc.analyticsRepo.GetMuscleVolumeDistribution(ctx, userID)
	if err != nil {
		if service.IsNoAnalyticsData(err) {
			return nil, domainerr.ErrNoAnalyticsData
		}
		return nil, fmt.Errorf("get muscle volume distribution: %w", err)
	}
	return mapper.MuscleVolumeDistributionToResultList(results), nil
}
