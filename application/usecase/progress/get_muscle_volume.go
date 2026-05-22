package progress

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetMuscleVolumeUseCase interface {
	GetMuscleVolume(ctx context.Context, userID string) ([]dto.MuscleVolumeResult, error)
}

type getMuscleVolumeUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewGetMuscleVolumeUseCase(progressRepo repository.ProgressRepository) GetMuscleVolumeUseCase {
	return &getMuscleVolumeUseCase{progressRepo: progressRepo}
}

func (uc *getMuscleVolumeUseCase) GetMuscleVolume(ctx context.Context, userID string) ([]dto.MuscleVolumeResult, error) {
	records, err := uc.progressRepo.FindMuscleVolume(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find muscle volume: %w", err)
	}

	return mapper.MuscleVolumeToResultList(records), nil
}
