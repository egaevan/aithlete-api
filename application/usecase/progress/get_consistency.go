package progress

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetConsistencyUseCase interface {
	GetConsistency(ctx context.Context, userID string) ([]dto.ConsistencyResult, error)
}

type getConsistencyUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewGetConsistencyUseCase(progressRepo repository.ProgressRepository) GetConsistencyUseCase {
	return &getConsistencyUseCase{progressRepo: progressRepo}
}

func (uc *getConsistencyUseCase) GetConsistency(ctx context.Context, userID string) ([]dto.ConsistencyResult, error) {
	records, err := uc.progressRepo.FindConsistency(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find consistency: %w", err)
	}

	return mapper.ConsistencyToResultList(records), nil
}
