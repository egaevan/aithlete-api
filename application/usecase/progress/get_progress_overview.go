package progress

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetProgressOverviewUseCase interface {
	GetOverview(ctx context.Context, userID string) (*dto.ProgressOverviewResult, error)
}

type getProgressOverviewUseCase struct {
	progressRepo repository.ProgressRepository
}

func NewGetProgressOverviewUseCase(progressRepo repository.ProgressRepository) GetProgressOverviewUseCase {
	return &getProgressOverviewUseCase{progressRepo: progressRepo}
}

func (uc *getProgressOverviewUseCase) GetOverview(ctx context.Context, userID string) (*dto.ProgressOverviewResult, error) {
	bwRecords, err := uc.progressRepo.FindBodyWeightByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find body weight: %w", err)
	}

	strengthRecords, err := uc.progressRepo.FindStrengthByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find strength: %w", err)
	}

	consistencyRecords, err := uc.progressRepo.FindConsistency(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find consistency: %w", err)
	}

	var totalVolume float64
	for _, sr := range strengthRecords {
		totalVolume += sr.Volume
	}

	var totalWorkouts, currentStreak, longestStreak int
	for _, c := range consistencyRecords {
		totalWorkouts += c.WorkoutsCompleted
		if c.Streak > longestStreak {
			longestStreak = c.Streak
		}
	}
	if len(consistencyRecords) > 0 {
		currentStreak = consistencyRecords[len(consistencyRecords)-1].Streak
	}

	var latestBW *dto.BodyWeightResult
	if len(bwRecords) > 0 {
		latestBW = mapper.BodyWeightToResult(&bwRecords[len(bwRecords)-1])
	}

	return &dto.ProgressOverviewResult{
		TotalWorkouts:    totalWorkouts,
		CurrentStreak:    currentStreak,
		LongestStreak:    longestStreak,
		TotalVolume:      totalVolume,
		WorkoutsThisWeek: 0,
		BodyWeight:       latestBW,
	}, nil
}
