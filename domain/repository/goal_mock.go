package repository

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockGoalRepository struct {
	Goals   map[string]*entity.Goal
	counter int
}

func NewMockGoalRepository() *MockGoalRepository {
	return &MockGoalRepository{Goals: make(map[string]*entity.Goal)}
}

func (m *MockGoalRepository) FindByID(_ context.Context, id string) (*entity.Goal, error) {
	g, ok := m.Goals[id]
	if !ok {
		return nil, domainerr.ErrGoalNotFound
	}
	return g, nil
}

func (m *MockGoalRepository) FindByUserID(_ context.Context, userID string) ([]entity.Goal, error) {
	var result []entity.Goal
	for _, g := range m.Goals {
		if g.UserID == userID {
			result = append(result, *g)
		}
	}
	return result, nil
}

func (m *MockGoalRepository) Create(_ context.Context, g *entity.Goal) error {
	m.counter++
	g.ID = fmt.Sprintf("goal-%d", m.counter)
	m.Goals[g.ID] = g
	return nil
}

func (m *MockGoalRepository) Update(_ context.Context, g *entity.Goal) error {
	if _, ok := m.Goals[g.ID]; !ok {
		return domainerr.ErrGoalNotFound
	}
	m.Goals[g.ID] = g
	return nil
}

func (m *MockGoalRepository) Delete(_ context.Context, id string) error {
	if _, ok := m.Goals[id]; !ok {
		return domainerr.ErrGoalNotFound
	}
	delete(m.Goals, id)
	return nil
}

var _ GoalRepository = (*MockGoalRepository)(nil)
