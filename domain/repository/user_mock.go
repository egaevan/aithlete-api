package repository

import (
	"context"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type MockUserRepository struct {
	Users map[string]*entity.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{Users: make(map[string]*entity.User)}
}

func (m *MockUserRepository) FindByID(_ context.Context, id string) (*entity.User, error) {
	for _, u := range m.Users {
		if u.GetID() == id {
			return u, nil
		}
	}
	return nil, domainerr.ErrUserNotFound
}

func (m *MockUserRepository) FindByEmail(_ context.Context, email string) (*entity.User, error) {
	for _, u := range m.Users {
		if u.GetEmail() == email {
			return u, nil
		}
	}
	return nil, domainerr.ErrUserNotFound
}

func (m *MockUserRepository) Create(_ context.Context, u *entity.User) error {
	u.SetID("user-" + u.GetEmail())
	m.Users[u.GetEmail()] = u
	return nil
}

func (m *MockUserRepository) Update(_ context.Context, u *entity.User) error {
	m.Users[u.GetEmail()] = u
	return nil
}
