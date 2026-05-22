package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*entity.User
	emails map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[string]*entity.User),
		emails: make(map[string]string),
	}
}

func (r *InMemoryUserRepository) FindByID(_ context.Context, id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("find user by id %s: %w", id, domainerr.ErrUserNotFound)
	}
	return u, nil
}

func (r *InMemoryUserRepository) FindByEmail(_ context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.emails[email]
	if !ok {
		return nil, fmt.Errorf("find user by email %s: %w", email, domainerr.ErrUserNotFound)
	}
	return r.users[id], nil
}

func (r *InMemoryUserRepository) Create(_ context.Context, u *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("user_%d", time.Now().UnixNano())
	u.SetID(id)
	r.users[id] = u
	r.emails[u.GetEmail()] = id
	return nil
}

func (r *InMemoryUserRepository) Update(_ context.Context, u *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := u.GetID()
	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("update user: %w", domainerr.ErrUserNotFound)
	}
	oldEmail := r.users[id].GetEmail()
	if oldEmail != u.GetEmail() {
		delete(r.emails, oldEmail)
		r.emails[u.GetEmail()] = id
	}
	r.users[id] = u
	return nil
}
