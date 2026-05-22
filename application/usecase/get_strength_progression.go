package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetStrengthProgressionUseCase interface {
	GetStrengthProgression(ctx context.Context, userID string) ([]dto.StrengthResult, error)
}

type getStrengthProgressionUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewGetStrengthProgressionUseCase(progressRepo repository.ProgressRepository) GetStrengthProgressionUseCase {
	return &getStrengthProgressionUseCase{progressRepo: progressRepo}
}

func (uc *getStrengthProgressionUseCase) GetStrengthProgression(ctx context.Context, userID string) ([]dto.StrengthResult, error) {
	records, err := uc.progressRepo.FindStrengthByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find strength: %w", err)
	}

	return mapper.StrengthToResultList(records), nil
}
