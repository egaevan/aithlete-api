package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/jackc/pgx/v5"
)

type ScheduleRepository struct {
	pool *Pool
}

func NewScheduleRepository(pool *Pool) *ScheduleRepository {
	return &ScheduleRepository{pool: pool}
}

func (r *ScheduleRepository) FindByID(ctx context.Context, id string) (*entity.Schedule, error) {
	var (
		userID, date, schedTime, title, duration, typ, notes string
		completed                                            bool
		createdAt, updatedAt                                 time.Time
	)

	err := r.pool.QueryRow(ctx, `
		SELECT id, user_id, date, time, title, duration, type, notes, completed, created_at, updated_at
		FROM schedules WHERE id = $1
	`, id).Scan(
		&id, &userID, &date, &schedTime, &title, &duration, &typ, &notes,
		&completed, &createdAt, &updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find schedule by id %s: %w", id, domainerr.ErrScheduleNotFound)
		}
		return nil, fmt.Errorf("find schedule by id: %w", err)
	}

	return &entity.Schedule{
		ID:        id,
		UserID:    userID,
		Date:      date,
		Time:      schedTime,
		Title:     title,
		Duration:  duration,
		Type:      typ,
		Notes:     notes,
		Completed: completed,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *ScheduleRepository) FindByUserID(ctx context.Context, userID string) ([]entity.Schedule, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, date, time, title, duration, type, notes, completed, created_at, updated_at
		FROM schedules WHERE user_id = $1
		ORDER BY date, time
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("find schedules by user id: %w", err)
	}
	defer rows.Close()

	var schedules []entity.Schedule
	for rows.Next() {
		var (
			id, scanUserID, date, schedTime, title, duration, typ, notes string
			completed                                                    bool
			createdAt, updatedAt                                         time.Time
		)
		if err := rows.Scan(
			&id, &scanUserID, &date, &schedTime, &title, &duration, &typ, &notes,
			&completed, &createdAt, &updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan schedule row: %w", err)
		}

		schedules = append(schedules, entity.Schedule{
			ID:        id,
			UserID:    scanUserID,
			Date:      date,
			Time:      schedTime,
			Title:     title,
			Duration:  duration,
			Type:      typ,
			Notes:     notes,
			Completed: completed,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return schedules, rows.Err()
}

func (r *ScheduleRepository) FindByUserIDAndDate(ctx context.Context, userID, date string) ([]entity.Schedule, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, date, time, title, duration, type, notes, completed, created_at, updated_at
		FROM schedules WHERE user_id = $1 AND date = $2
		ORDER BY time
	`, userID, date)
	if err != nil {
		return nil, fmt.Errorf("find schedules by user id and date: %w", err)
	}
	defer rows.Close()

	var schedules []entity.Schedule
	for rows.Next() {
		var (
			id, scanUserID, scanDate, schedTime, title, duration, typ, notes string
			completed                                                        bool
			createdAt, updatedAt                                             time.Time
		)
		if err := rows.Scan(
			&id, &scanUserID, &scanDate, &schedTime, &title, &duration, &typ, &notes,
			&completed, &createdAt, &updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan schedule row: %w", err)
		}

		schedules = append(schedules, entity.Schedule{
			ID:        id,
			UserID:    scanUserID,
			Date:      scanDate,
			Time:      schedTime,
			Title:     title,
			Duration:  duration,
			Type:      typ,
			Notes:     notes,
			Completed: completed,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return schedules, rows.Err()
}

func (r *ScheduleRepository) Create(ctx context.Context, s *entity.Schedule) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO schedules (user_id, date, time, title, duration, type, notes, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`, s.UserID, s.Date, s.Time, s.Title, s.Duration, s.Type, s.Notes, s.Completed, s.CreatedAt, s.UpdatedAt).Scan(&id)
	if err != nil {
		return fmt.Errorf("create schedule: %w", err)
	}

	s.ID = id
	return nil
}

func (r *ScheduleRepository) Update(ctx context.Context, s *entity.Schedule) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE schedules SET user_id=$1, date=$2, time=$3, title=$4, duration=$5,
		                     type=$6, notes=$7, completed=$8, updated_at=$9
		WHERE id=$10
	`, s.UserID, s.Date, s.Time, s.Title, s.Duration, s.Type, s.Notes, s.Completed, s.UpdatedAt, s.ID)
	if err != nil {
		return fmt.Errorf("update schedule: %w", err)
	}

	return nil
}

func (r *ScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM schedules WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete schedule: %w", err)
	}

	return nil
}
