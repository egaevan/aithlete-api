package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetBodyWeightHistoryUseCase interface {
	GetBodyWeightHistory(ctx context.Context, userID string) ([]dto.BodyWeightResult, error)
}

type getBodyWeightHistoryUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewGetBodyWeightHistoryUseCase(progressRepo repository.ProgressRepository) GetBodyWeightHistoryUseCase {
	return &getBodyWeightHistoryUseCase{progressRepo: progressRepo}
}

func (uc *getBodyWeightHistoryUseCase) GetBodyWeightHistory(ctx context.Context, userID string) ([]dto.BodyWeightResult, error) {
	records, err := uc.progressRepo.FindBodyWeightByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find body weight: %w", err)
	}

	return mapper.BodyWeightToResultList(records), nil
}
