package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type AddBodyWeightUseCase interface {
	Add(ctx context.Context, userID, date string, weight, bodyFatPercentage float64) (*dto.BodyWeightResult, error)
}

type addBodyWeightUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewAddBodyWeightUseCase(progressRepo repository.ProgressRepository) AddBodyWeightUseCase {
	return &addBodyWeightUseCase{progressRepo: progressRepo}
}

func (uc *addBodyWeightUseCase) Add(ctx context.Context, userID, date string, weight, bodyFatPercentage float64) (*dto.BodyWeightResult, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	bw := entity.NewBodyWeight(userID, date, weight, bodyFatPercentage)
	if err := uc.progressRepo.AddBodyWeight(ctx, bw); err != nil {
		return nil, fmt.Errorf("add body weight: %w", err)
	}
	return mapper.BodyWeightToResult(bw), nil
}
