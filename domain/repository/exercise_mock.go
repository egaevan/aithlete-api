package repository

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockExerciseRepository struct {
	Exercises map[string]*entity.Exercise
	counter   int
}

func NewMockExerciseRepository() *MockExerciseRepository {
	return &MockExerciseRepository{Exercises: make(map[string]*entity.Exercise)}
}

func (m *MockExerciseRepository) FindAll(_ context.Context) ([]entity.Exercise, error) {
	var result []entity.Exercise
	for _, e := range m.Exercises {
		result = append(result, *e)
	}
	return result, nil
}

func (m *MockExerciseRepository) FindByID(_ context.Context, id string) (*entity.Exercise, error) {
	e, ok := m.Exercises[id]
	if !ok {
		return nil, domainerr.ErrExerciseNotFound
	}
	return e, nil
}

func (m *MockExerciseRepository) FindByMuscleGroup(_ context.Context, muscleGroup string) ([]entity.Exercise, error) {
	var result []entity.Exercise
	for _, e := range m.Exercises {
		if e.MuscleGroup == muscleGroup {
			result = append(result, *e)
		}
	}
	return result, nil
}

func (m *MockExerciseRepository) Create(_ context.Context, e *entity.Exercise) error {
	m.counter++
	e.ID = fmt.Sprintf("exercise-%d", m.counter)
	m.Exercises[e.ID] = e
	return nil
}

var _ ExerciseRepository = (*MockExerciseRepository)(nil)
