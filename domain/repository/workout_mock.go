package repository

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockWorkoutRepository struct {
	Workouts  map[string]*entity.Workout
	counter   int
}

func NewMockWorkoutRepository() *MockWorkoutRepository {
	return &MockWorkoutRepository{Workouts: make(map[string]*entity.Workout)}
}

func (m *MockWorkoutRepository) FindByID(_ context.Context, id string) (*entity.Workout, error) {
	w, ok := m.Workouts[id]
	if !ok {
		return nil, domainerr.ErrWorkoutNotFound
	}
	return w, nil
}

func (m *MockWorkoutRepository) FindByUserID(_ context.Context, userID string) ([]entity.Workout, error) {
	var result []entity.Workout
	for _, w := range m.Workouts {
		if w.UserID == userID {
			result = append(result, *w)
		}
	}
	return result, nil
}

func (m *MockWorkoutRepository) Create(_ context.Context, w *entity.Workout) error {
	m.counter++
	w.ID = fmt.Sprintf("workout-%d", m.counter)
	m.Workouts[w.ID] = w
	return nil
}

func (m *MockWorkoutRepository) Update(_ context.Context, w *entity.Workout) error {
	if _, ok := m.Workouts[w.ID]; !ok {
		return domainerr.ErrWorkoutNotFound
	}
	m.Workouts[w.ID] = w
	return nil
}

func (m *MockWorkoutRepository) Delete(_ context.Context, id string) error {
	if _, ok := m.Workouts[id]; !ok {
		return domainerr.ErrWorkoutNotFound
	}
	delete(m.Workouts, id)
	return nil
}
