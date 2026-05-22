package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/jackc/pgx/v5"
)

type GoalRepository struct {
	pool *Pool
}

func NewGoalRepository(pool *Pool) *GoalRepository {
	return &GoalRepository{pool: pool}
}

func (r *GoalRepository) FindByID(ctx context.Context, id string) (*entity.Goal, error) {
	var (
		userID, title, typ, unit, period, deadline string
		target, current                            int
		completed                                  bool
		createdAt, updatedAt                       time.Time
	)

	err := r.pool.QueryRow(ctx, `
		SELECT id, user_id, title, type, target, current, unit, period, deadline, completed, created_at, updated_at
		FROM goals WHERE id = $1
	`, id).Scan(
		&id, &userID, &title, &typ, &target, &current, &unit, &period, &deadline,
		&completed, &createdAt, &updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find goal by id %s: %w", id, domainerr.ErrGoalNotFound)
		}
		return nil, fmt.Errorf("find goal by id: %w", err)
	}

	return &entity.Goal{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Type:      typ,
		Target:    target,
		Current:   current,
		Unit:      unit,
		Period:    period,
		Deadline:  deadline,
		Completed: completed,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *GoalRepository) FindByUserID(ctx context.Context, userID string) ([]entity.Goal, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, title, type, target, current, unit, period, deadline, completed, created_at, updated_at
		FROM goals WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("find goals by user id: %w", err)
	}
	defer rows.Close()

	var goals []entity.Goal
	for rows.Next() {
		var (
			id, scanUserID, title, typ, unit, period, deadline string
			target, current                                    int
			completed                                          bool
			createdAt, updatedAt                               time.Time
		)
		if err := rows.Scan(
			&id, &scanUserID, &title, &typ, &target, &current, &unit, &period, &deadline,
			&completed, &createdAt, &updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan goal row: %w", err)
		}

		goals = append(goals, entity.Goal{
			ID:        id,
			UserID:    scanUserID,
			Title:     title,
			Type:      typ,
			Target:    target,
			Current:   current,
			Unit:      unit,
			Period:    period,
			Deadline:  deadline,
			Completed: completed,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return goals, rows.Err()
}

func (r *GoalRepository) Create(ctx context.Context, g *entity.Goal) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO goals (user_id, title, type, target, current, unit, period, deadline, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`, g.UserID, g.Title, g.Type, g.Target, g.Current, g.Unit, g.Period, g.Deadline, g.Completed, g.CreatedAt, g.UpdatedAt).Scan(&id)
	if err != nil {
		return fmt.Errorf("create goal: %w", err)
	}

	g.ID = id
	return nil
}

func (r *GoalRepository) Update(ctx context.Context, g *entity.Goal) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE goals SET user_id=$1, title=$2, type=$3, target=$4, current=$5,
		                 unit=$6, period=$7, deadline=$8, completed=$9, updated_at=$10
		WHERE id=$11
	`, g.UserID, g.Title, g.Type, g.Target, g.Current, g.Unit, g.Period, g.Deadline, g.Completed, g.UpdatedAt, g.ID)
	if err != nil {
		return fmt.Errorf("update goal: %w", err)
	}

	return nil
}

func (r *GoalRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM goals WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete goal: %w", err)
	}

	return nil
}
