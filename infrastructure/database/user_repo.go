package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	pool *Pool
}

func NewUserRepository(pool *Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var (
		email, name, password, avatar, birthday, gender string
		createdAt, updatedAt                            time.Time
	)
	err := r.pool.QueryRow(ctx, `
		SELECT id, email, name, password, avatar, birthday, gender, created_at, updated_at
		FROM users WHERE id = $1
	`, id).Scan(
		&id, &email, &name, &password,
		&avatar, &birthday, &gender,
		&createdAt, &updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find user by id %s: %w", id, domainerr.ErrUserNotFound)
		}
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	return entity.RebuildUser(id, email, name, password, avatar, birthday, gender, createdAt, updatedAt), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var (
		id, name, password, avatar, birthday, gender string
		createdAt, updatedAt                         time.Time
	)
	err := r.pool.QueryRow(ctx, `
		SELECT id, email, name, password, avatar, birthday, gender, created_at, updated_at
		FROM users WHERE email = $1
	`, email).Scan(
		&id, &email, &name, &password,
		&avatar, &birthday, &gender,
		&createdAt, &updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find user by email %s: %w", email, domainerr.ErrUserNotFound)
		}
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	return entity.RebuildUser(id, email, name, password, avatar, birthday, gender, createdAt, updatedAt), nil
}

func (r *UserRepository) Create(ctx context.Context, u *entity.User) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO users (email, name, password, avatar, birthday, gender, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, u.GetEmail(), u.GetName(), u.PasswordHash(), u.GetAvatar(), u.GetBirthday(), u.GetGender(), u.GetCreatedAt(), u.GetUpdatedAt()).Scan(&id)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	u.SetID(id)
	return nil
}

func (r *UserRepository) Update(ctx context.Context, u *entity.User) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE users SET email=$1, name=$2, password=$3, avatar=$4,
		                 birthday=$5, gender=$6, updated_at=$7
		WHERE id=$8
	`, u.GetEmail(), u.GetName(), u.PasswordHash(), u.GetAvatar(), u.GetBirthday(), u.GetGender(), u.GetUpdatedAt(), u.GetID())
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}
