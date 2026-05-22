package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/aithlete/aithlete-api/domain/entity"
)

type MockProgressRepository struct {
	BodyWeights      map[string]*entity.BodyWeight
	StrengthRecords  map[string]*entity.StrengthRecord
	ConsistencyData  map[string]*entity.Consistency
	MuscleVolumeData map[string]*entity.MuscleVolume
	bwCounter        int
	srCounter        int
}

func NewMockProgressRepository() *MockProgressRepository {
	return &MockProgressRepository{
		BodyWeights:      make(map[string]*entity.BodyWeight),
		StrengthRecords:  make(map[string]*entity.StrengthRecord),
		ConsistencyData:  make(map[string]*entity.Consistency),
		MuscleVolumeData: make(map[string]*entity.MuscleVolume),
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

func (m *MockProgressRepository) AddStrengthRecord(_ context.Context, sr *entity.StrengthRecord) error {
	m.srCounter++
	sr.ID = fmt.Sprintf("sr-%d", m.srCounter)
	m.StrengthRecords[sr.ID] = sr
	return nil
}

func (m *MockProgressRepository) FindConsistency(_ context.Context, userID string) ([]entity.Consistency, error) {
	var result []entity.Consistency
	for _, c := range m.ConsistencyData {
		result = append(result, *c)
	}
	return result, nil
}

func (m *MockProgressRepository) UpsertConsistency(_ context.Context, userID, week string, completed int) error {
	key := userID + ":" + week
	existing, ok := m.ConsistencyData[key]
	if ok {
		existing.WorkoutsCompleted += completed
		return nil
	}

	newStreak := 1
	for k, pc := range m.ConsistencyData {
		uid, w, ok := strings.Cut(k, ":")
		if ok && uid == userID && weeksAreConsecutive(w, week) {
			newStreak = pc.Streak + completed
			break
		}
	}

	m.ConsistencyData[key] = &entity.Consistency{
		Week:              week,
		WorkoutsCompleted: completed,
		WorkoutsPlanned:   0,
		Streak:            newStreak,
	}
	return nil
}

func weeksAreConsecutive(a, b string) bool {
	var ay, aw, by, bw int
	if _, err := fmt.Sscanf(a, "%d-W%d", &ay, &aw); err != nil {
		return false
	}
	if _, err := fmt.Sscanf(b, "%d-W%d", &by, &bw); err != nil {
		return false
	}

	if by == ay && bw == aw+1 {
		return true
	}
	if by == ay+1 && aw >= 52 && bw == 1 {
		return true
	}
	return false
}

func (m *MockProgressRepository) FindMuscleVolume(_ context.Context, userID string) ([]entity.MuscleVolume, error) {
	var result []entity.MuscleVolume
	for _, mv := range m.MuscleVolumeData {
		result = append(result, *mv)
	}
	return result, nil
}

func (m *MockProgressRepository) UpsertMuscleVolume(_ context.Context, userID, muscleGroup string, volume float64) error {
	key := userID + ":" + muscleGroup
	existing, ok := m.MuscleVolumeData[key]
	if ok {
		existing.Volume += volume
		existing.Sessions++
	} else {
		m.MuscleVolumeData[key] = &entity.MuscleVolume{
			MuscleGroup: muscleGroup,
			Volume:      volume,
			Sessions:    1,
			Trend:       "up",
		}
	}
	return nil
}

var _ ProgressRepository = (*MockProgressRepository)(nil)
