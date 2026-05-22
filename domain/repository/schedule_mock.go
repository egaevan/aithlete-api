package repository

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockScheduleRepository struct {
	Schedules map[string]*entity.Schedule
	counter   int
}

func NewMockScheduleRepository() *MockScheduleRepository {
	return &MockScheduleRepository{Schedules: make(map[string]*entity.Schedule)}
}

func (m *MockScheduleRepository) FindByID(_ context.Context, id string) (*entity.Schedule, error) {
	s, ok := m.Schedules[id]
	if !ok {
		return nil, domainerr.ErrScheduleNotFound
	}
	return s, nil
}

func (m *MockScheduleRepository) FindByUserID(_ context.Context, userID string) ([]entity.Schedule, error) {
	var result []entity.Schedule
	for _, s := range m.Schedules {
		if s.UserID == userID {
			result = append(result, *s)
		}
	}
	return result, nil
}

func (m *MockScheduleRepository) FindByUserIDAndDate(_ context.Context, userID, date string) ([]entity.Schedule, error) {
	var result []entity.Schedule
	for _, s := range m.Schedules {
		if s.UserID == userID && s.Date == date {
			result = append(result, *s)
		}
	}
	return result, nil
}

func (m *MockScheduleRepository) Create(_ context.Context, s *entity.Schedule) error {
	m.counter++
	s.ID = fmt.Sprintf("schedule-%d", m.counter)
	m.Schedules[s.ID] = s
	return nil
}

func (m *MockScheduleRepository) Update(_ context.Context, s *entity.Schedule) error {
	if _, ok := m.Schedules[s.ID]; !ok {
		return domainerr.ErrScheduleNotFound
	}
	m.Schedules[s.ID] = s
	return nil
}

func (m *MockScheduleRepository) Delete(_ context.Context, id string) error {
	if _, ok := m.Schedules[id]; !ok {
		return domainerr.ErrScheduleNotFound
	}
	delete(m.Schedules, id)
	return nil
}
