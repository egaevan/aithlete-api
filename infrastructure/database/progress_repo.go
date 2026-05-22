package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/jackc/pgx/v5"
)

type ProgressRepository struct {
	pool *Pool
}

func NewProgressRepository(pool *Pool) *ProgressRepository {
	return &ProgressRepository{pool: pool}
}

func (r *ProgressRepository) FindBodyWeightByUserID(ctx context.Context, userID string) ([]entity.BodyWeight, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, date, weight, body_fat_percentage, created_at
		FROM body_weight WHERE user_id = $1
		ORDER BY date DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("find body weight: %w", err)
	}
	defer rows.Close()

	var records []entity.BodyWeight
	for rows.Next() {
		var (
			id, scanUserID, date string
			weight, bodyFat      float64
			createdAt            time.Time
		)
		if err := rows.Scan(&id, &scanUserID, &date, &weight, &bodyFat, &createdAt); err != nil {
			return nil, fmt.Errorf("scan body weight: %w", err)
		}
		records = append(records, entity.BodyWeight{
			ID: id, UserID: scanUserID, Date: date,
			Weight: weight, BodyFatPercentage: bodyFat, CreatedAt: createdAt,
		})
	}
	return records, rows.Err()
}

func (r *ProgressRepository) AddBodyWeight(ctx context.Context, bw *entity.BodyWeight) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO body_weight (user_id, date, weight, body_fat_percentage, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, bw.UserID, bw.Date, bw.Weight, bw.BodyFatPercentage, bw.CreatedAt).Scan(&id)
	if err != nil {
		return fmt.Errorf("add body weight: %w", err)
	}
	bw.ID = id
	return nil
}

func (r *ProgressRepository) FindStrengthByUserID(ctx context.Context, userID string) ([]entity.StrengthRecord, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, exercise, date, one_rep_max, volume
		FROM strength_progression WHERE user_id = $1
		ORDER BY date DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("find strength: %w", err)
	}
	defer rows.Close()

	var records []entity.StrengthRecord
	for rows.Next() {
		var (
			id, scanUserID, exercise, date string
			oneRepMax, volume              float64
		)
		if err := rows.Scan(&id, &scanUserID, &exercise, &date, &oneRepMax, &volume); err != nil {
			return nil, fmt.Errorf("scan strength: %w", err)
		}
		records = append(records, entity.StrengthRecord{
			ID: id, UserID: scanUserID, Exercise: exercise, Date: date,
			OneRepMax: oneRepMax, Volume: volume,
		})
	}
	return records, rows.Err()
}

func (r *ProgressRepository) AddStrengthRecord(ctx context.Context, sr *entity.StrengthRecord) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO strength_progression (user_id, exercise, date, one_rep_max, volume)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, sr.UserID, sr.Exercise, sr.Date, sr.OneRepMax, sr.Volume).Scan(&id)
	if err != nil {
		return fmt.Errorf("add strength record: %w", err)
	}
	sr.ID = id
	return nil
}

func (r *ProgressRepository) UpsertConsistency(ctx context.Context, userID, week string, completed int) error {
	newStreak := 1

	var prevWeek string
	var prevStreak int
	err := r.pool.QueryRow(ctx, `
		SELECT week, streak
		FROM consistency
		WHERE user_id = $1 AND week < $2
		ORDER BY week DESC
		LIMIT 1
	`, userID, week).Scan(&prevWeek, &prevStreak)
	if err == nil && weeksAreConsecutive(prevWeek, week) {
		newStreak = prevStreak + completed
	}

	_, err = r.pool.Exec(ctx, `
		INSERT INTO consistency (user_id, week, workouts_completed, workouts_planned, streak)
		VALUES ($1, $2, $3, 0, $4)
		ON CONFLICT ON CONSTRAINT consistency_user_week_idx
		DO UPDATE SET workouts_completed = consistency.workouts_completed + $3,
		              streak = GREATEST(consistency.streak, $4)
	`, userID, week, completed, newStreak)
	if err != nil {
		return fmt.Errorf("upsert consistency: %w", err)
	}
	return nil
}

func weeksAreConsecutive(a, b string) bool {
	var ay, aw, by, bw int
	if _, err := fmt.Sscanf(a, "%d-W%d", &ay, &aw); err != nil {
		return false
	}
	if _, err := fmt.Sscanf(b, "%d-W%d", &by, &bw); err != nil {
		return false
	}

	if by == ay && bw == aw+1 {
		return true
	}
	if by == ay+1 && aw >= 52 && bw == 1 {
		return true
	}
	return false
}

func (r *ProgressRepository) UpsertMuscleVolume(ctx context.Context, userID, muscleGroup string, volume float64) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO muscle_volume (user_id, muscle_group, volume, sessions, trend)
		VALUES ($1, $2, $3, 1, 'up')
		ON CONFLICT ON CONSTRAINT muscle_volume_user_muscle_idx
		DO UPDATE SET volume = muscle_volume.volume + $3,
		              sessions = muscle_volume.sessions + 1
	`, userID, muscleGroup, volume)
	if err != nil {
		return fmt.Errorf("upsert muscle volume: %w", err)
	}
	return nil
}

func (r *ProgressRepository) FindConsistency(ctx context.Context, userID string) ([]entity.Consistency, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT week, workouts_completed, workouts_planned, streak
		FROM consistency WHERE user_id = $1
		ORDER BY week DESC
	`, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("find consistency: %w", err)
	}
	defer rows.Close()

	var records []entity.Consistency
	for rows.Next() {
		var week string
		var completed, planned, streak int
		if err := rows.Scan(&week, &completed, &planned, &streak); err != nil {
			return nil, fmt.Errorf("scan consistency: %w", err)
		}
		records = append(records, entity.Consistency{
			Week: week, WorkoutsCompleted: completed,
			WorkoutsPlanned: planned, Streak: streak,
		})
	}
	return records, rows.Err()
}

func (r *ProgressRepository) FindMuscleVolume(ctx context.Context, userID string) ([]entity.MuscleVolume, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT muscle_group, volume, sessions, trend
		FROM muscle_volume WHERE user_id = $1
		ORDER BY volume DESC
	`, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("find muscle volume: %w", err)
	}
	defer rows.Close()

	var records []entity.MuscleVolume
	for rows.Next() {
		var group, trend string
		var volume float64
		var sessions int
		if err := rows.Scan(&group, &volume, &sessions, &trend); err != nil {
			return nil, fmt.Errorf("scan muscle volume: %w", err)
		}
		records = append(records, entity.MuscleVolume{
			MuscleGroup: group, Volume: volume,
			Sessions: sessions, Trend: trend,
		})
	}
	return records, rows.Err()
}
