package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/jackc/pgx/v5"
)

type WorkoutRepository struct {
	pool *Pool
}

func NewWorkoutRepository(pool *Pool) *WorkoutRepository {
	return &WorkoutRepository{pool: pool}
}

func (r *WorkoutRepository) FindByID(ctx context.Context, id string) (*entity.Workout, error) {
	var (
		userID, name, date, weightUnit, notes string
		duration, calories, avgHeartRate      int
		completed                             bool
		exercisesJSON                         []byte
		createdAt, updatedAt                  time.Time
	)

	err := r.pool.QueryRow(ctx, `
		SELECT id, user_id, name, date, duration, weight_unit, notes,
		       completed, calories, avg_heart_rate, exercises, created_at, updated_at
		FROM workouts WHERE id = $1
	`, id).Scan(
		&id, &userID, &name, &date, &duration, &weightUnit, &notes,
		&completed, &calories, &avgHeartRate, &exercisesJSON, &createdAt, &updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find workout by id %s: %w", id, domainerr.ErrWorkoutNotFound)
		}
		return nil, fmt.Errorf("find workout by id: %w", err)
	}

	var exercises []entity.WorkoutExercise
	if len(exercisesJSON) > 0 {
		if err := json.Unmarshal(exercisesJSON, &exercises); err != nil {
			return nil, fmt.Errorf("unmarshal exercises: %w", err)
		}
	}

	return &entity.Workout{
		ID:           id,
		UserID:       userID,
		Name:         name,
		Date:         date,
		Duration:     duration,
		WeightUnit:   weightUnit,
		Notes:        notes,
		Completed:    completed,
		Calories:     calories,
		AvgHeartRate: avgHeartRate,
		Exercises:    exercises,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func (r *WorkoutRepository) FindByUserID(ctx context.Context, userID string) ([]entity.Workout, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, name, date, duration, weight_unit, notes,
		       completed, calories, avg_heart_rate, exercises, created_at, updated_at
		FROM workouts WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("find workouts by user id: %w", err)
	}
	defer rows.Close()

	var workouts []entity.Workout
	for rows.Next() {
		var (
			id, name, date, weightUnit, notes    string
			duration, calories, avgHeartRate      int
			completed                             bool
			exercisesJSON                         []byte
			createdAt, updatedAt                  time.Time
			scanUserID                            string
		)
		if err := rows.Scan(
			&id, &scanUserID, &name, &date, &duration, &weightUnit, &notes,
			&completed, &calories, &avgHeartRate, &exercisesJSON, &createdAt, &updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan workout row: %w", err)
		}

		var exercises []entity.WorkoutExercise
		if len(exercisesJSON) > 0 {
			if err := json.Unmarshal(exercisesJSON, &exercises); err != nil {
				return nil, fmt.Errorf("unmarshal exercises: %w", err)
			}
		}

		workouts = append(workouts, entity.Workout{
			ID:           id,
			UserID:       scanUserID,
			Name:         name,
			Date:         date,
			Duration:     duration,
			WeightUnit:   weightUnit,
			Notes:        notes,
			Completed:    completed,
			Calories:     calories,
			AvgHeartRate: avgHeartRate,
			Exercises:    exercises,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		})
	}

	return workouts, rows.Err()
}

func (r *WorkoutRepository) Create(ctx context.Context, w *entity.Workout) error {
	exercisesJSON, err := json.Marshal(w.Exercises)
	if err != nil {
		return fmt.Errorf("marshal exercises: %w", err)
	}

	var id string
	err = r.pool.QueryRow(ctx, `
		INSERT INTO workouts (user_id, name, date, duration, weight_unit, notes,
		                      completed, calories, avg_heart_rate, exercises, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`, w.UserID, w.Name, w.Date, w.Duration, w.WeightUnit, w.Notes,
		w.Completed, w.Calories, w.AvgHeartRate, exercisesJSON, w.CreatedAt, w.UpdatedAt).Scan(&id)
	if err != nil {
		return fmt.Errorf("create workout: %w", err)
	}

	w.ID = id
	return nil
}

func (r *WorkoutRepository) Update(ctx context.Context, w *entity.Workout) error {
	exercisesJSON, err := json.Marshal(w.Exercises)
	if err != nil {
		return fmt.Errorf("marshal exercises: %w", err)
	}

	_, err = r.pool.Exec(ctx, `
		UPDATE workouts SET user_id=$1, name=$2, date=$3, duration=$4, weight_unit=$5,
		                    notes=$6, completed=$7, calories=$8, avg_heart_rate=$9,
		                    exercises=$10, updated_at=$11
		WHERE id=$12
	`, w.UserID, w.Name, w.Date, w.Duration, w.WeightUnit, w.Notes,
		w.Completed, w.Calories, w.AvgHeartRate, exercisesJSON, w.UpdatedAt, w.ID)
	if err != nil {
		return fmt.Errorf("update workout: %w", err)
	}

	return nil
}

func (r *WorkoutRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM workouts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete workout: %w", err)
	}

	return nil
}
