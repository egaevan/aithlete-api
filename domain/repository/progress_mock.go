package repository

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type MockProgressRepository struct {
	BodyWeights     map[string]*entity.BodyWeight
	StrengthRecords map[string]*entity.StrengthRecord
	bwCounter       int
	srCounter       int
}

func NewMockProgressRepository() *MockProgressRepository {
	return &MockProgressRepository{
		BodyWeights:     make(map[string]*entity.BodyWeight),
		StrengthRecords: make(map[string]*entity.StrengthRecord),
	}
}

func (m *MockProgressRepository) FindBodyWeightByUserID(_ context.Context, userID string) ([]entity.BodyWeight, error) {
	var result []entity.BodyWeight
	for _, bw := range m.BodyWeights {
		if bw.UserID == userID {
			result = append(result, *bw)
		}
	}
	return result, nil
}

func (m *MockProgressRepository) AddBodyWeight(_ context.Context, bw *entity.BodyWeight) error {
	m.bwCounter++
	bw.ID = fmt.Sprintf("bw-%d", m.bwCounter)
	m.BodyWeights[bw.ID] = bw
	return nil
}

func (m *MockProgressRepository) FindStrengthByUserID(_ context.Context, userID string) ([]entity.StrengthRecord, error) {
	var result []entity.StrengthRecord
	for _, sr := range m.StrengthRecords {
		if sr.UserID == userID {
			result = append(result, *sr)
		}
	}
	return result, nil
}

func (m *MockProgressRepository) FindConsistency(_ context.Context, userID string) ([]entity.Consistency, error) {
	return nil, nil
}

func (m *MockProgressRepository) FindMuscleVolume(_ context.Context, userID string) ([]entity.MuscleVolume, error) {
	return nil, nil
}

var _ ProgressRepository = (*MockProgressRepository)(nil)
