package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type AnalyticsRepository struct {
	pool *Pool
}

func NewAnalyticsRepository(pool *Pool) *AnalyticsRepository {
	return &AnalyticsRepository{pool: pool}
}

func (r *AnalyticsRepository) GetDashboard(ctx context.Context, userID string) (*entity.Dashboard, error) {
	now := time.Now()
	weekStart := now.AddDate(0, 0, -int(now.Weekday()))
	weekStartStr := weekStart.Format("2006-01-02")

	var caloriesBurned, activeMinutes int
	err := r.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(calories), 0), COALESCE(SUM(duration), 0)
		FROM workouts
		WHERE user_id = $1 AND completed = true AND date >= $2
	`, userID, weekStartStr).Scan(&caloriesBurned, &activeMinutes)
	if err != nil {
		return nil, fmt.Errorf("query dashboard stats: %w", err)
	}

	var goalsCompleted, goalsTotal int
	err = r.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(CASE WHEN completed THEN 1 ELSE 0 END), 0), COUNT(*)
		FROM goals WHERE user_id = $1
	`, userID).Scan(&goalsCompleted, &goalsTotal)
	if err != nil {
		return nil, fmt.Errorf("query dashboard goals: %w", err)
	}

	weeklyProgress, err := r.getWeeklyProgress(ctx, userID, now)
	if err != nil {
		return nil, fmt.Errorf("query weekly progress: %w", err)
	}

	streak, err := r.GetStreak(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("query streak: %w", err)
	}

	return &entity.Dashboard{
		Stats: entity.DashboardStats{
			CaloriesBurned:     caloriesBurned,
			CaloriesTrend:      "0%",
			ActiveMinutes:      activeMinutes,
			ActiveMinutesTrend: "0%",
			GoalsCompleted:     goalsCompleted,
			GoalsTotal:         goalsTotal,
			GoalsTrend:         "on track",
			AvgHeartRate:       0,
			HeartRateTrend:     "0%",
		},
		WeeklyProgress: weeklyProgress,
		Streak:         *streak,
	}, nil
}

func (r *AnalyticsRepository) GetWeeklyProgress(ctx context.Context, userID string) ([]entity.WeeklyProgressDay, error) {
	return r.getWeeklyProgress(ctx, userID, time.Now())
}

func (r *AnalyticsRepository) getWeeklyProgress(ctx context.Context, userID string, now time.Time) ([]entity.WeeklyProgressDay, error) {
	weekStart := now.AddDate(0, 0, -int(now.Weekday()))
	weekStartStr := weekStart.Format("2006-01-02")

	rows, err := r.pool.Query(ctx, `
		SELECT
			to_char(date::date, 'Day') AS day,
			COALESCE(SUM(calories), 0) AS calories,
			COALESCE(SUM(duration), 0) AS duration
		FROM workouts
		WHERE user_id = $1 AND completed = true AND date >= $2
		GROUP BY date::date
		ORDER BY date::date
	`, userID, weekStartStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.WeeklyProgressDay
	for rows.Next() {
		var d entity.WeeklyProgressDay
		if err := rows.Scan(&d.Day, &d.Calories, &d.Duration); err != nil {
			return nil, fmt.Errorf("scan weekly progress: %w", err)
		}
		results = append(results, d)
	}

	if results == nil {
		results = []entity.WeeklyProgressDay{}
	}
	return results, nil
}

func (r *AnalyticsRepository) GetStreak(ctx context.Context, userID string) (*entity.Streak, error) {
	var current int
	err := r.pool.QueryRow(ctx, `
		SELECT COALESCE(MAX(streak), 0)
		FROM consistency
		WHERE user_id = $1
	`, userID).Scan(&current)
	if err != nil {
		return nil, domainerr.ErrNoAnalyticsData
	}

	var personalBest int
	err = r.pool.QueryRow(ctx, `
		SELECT COALESCE(MAX(streak), 0)
		FROM consistency
		WHERE user_id = $1
	`, userID).Scan(&personalBest)
	if err != nil {
		personalBest = current
	}

	return &entity.Streak{
		Current:      current,
		PersonalBest: personalBest,
	}, nil
}

func (r *AnalyticsRepository) GetWeeklyVolume(ctx context.Context, userID string) ([]entity.WeeklyVolume, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			to_char(sp.date::date, 'IYYY"-W"IW') AS week,
			COALESCE(SUM(sp.volume), 0) AS volume
		FROM strength_progression sp
		WHERE sp.user_id = $1
		GROUP BY week
		ORDER BY week DESC
		LIMIT 12
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("query weekly volume: %w", err)
	}
	defer rows.Close()

	var results []entity.WeeklyVolume
	for rows.Next() {
		var v entity.WeeklyVolume
		if err := rows.Scan(&v.Week, &v.Volume); err != nil {
			return nil, fmt.Errorf("scan weekly volume: %w", err)
		}
		results = append(results, v)
	}

	if results == nil {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return results, nil
}

func (r *AnalyticsRepository) GetMuscleVolumeDistribution(ctx context.Context, userID string) ([]entity.MuscleVolumeDistribution, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT muscle_group, COALESCE(SUM(volume), 0) AS volume
		FROM muscle_volume
		WHERE user_id = $1
		GROUP BY muscle_group
		ORDER BY volume DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("query muscle volume: %w", err)
	}
	defer rows.Close()

	var results []entity.MuscleVolumeDistribution
	for rows.Next() {
		var m entity.MuscleVolumeDistribution
		if err := rows.Scan(&m.Muscle, &m.Volume); err != nil {
			return nil, fmt.Errorf("scan muscle volume: %w", err)
		}
		results = append(results, m)
	}

	if results == nil {
		return nil, domainerr.ErrNoAnalyticsData
	}
	return results, nil
}

func (r *AnalyticsRepository) GetOverview(ctx context.Context, userID string) (*entity.AnalyticsOverview, error) {
	var totalVolume int
	err := r.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(sp.volume), 0)
		FROM strength_progression sp
		WHERE sp.user_id = $1
	`, userID).Scan(&totalVolume)
	if err != nil {
		return nil, fmt.Errorf("query total volume: %w", err)
	}

	var avgSession, sessionsPerMonth int
	err = r.pool.QueryRow(ctx, `
		SELECT
			COALESCE(ROUND(AVG(duration)), 0),
			COUNT(*)
		FROM workouts
		WHERE user_id = $1 AND completed = true
		  AND date >= to_char(NOW() - INTERVAL '30 days', 'YYYY-MM-DD')
	`, userID).Scan(&avgSession, &sessionsPerMonth)
	if err != nil {
		return nil, fmt.Errorf("query sessions: %w", err)
	}

	var goalCompletion int
	err = r.pool.QueryRow(ctx, `
		SELECT COALESCE(ROUND(AVG(CASE WHEN completed THEN 100.0 ELSE 0 END)), 0)
		FROM goals WHERE user_id = $1
	`, userID).Scan(&goalCompletion)
	if err != nil {
		return nil, fmt.Errorf("query goal completion: %w", err)
	}

	return &entity.AnalyticsOverview{
		TotalVolume:           totalVolume,
		TotalVolumeTrend:      "0%",
		AvgSession:            avgSession,
		AvgSessionTrend:       "0 min",
		SessionsPerMonth:      sessionsPerMonth,
		SessionsPerMonthTrend: "0",
		GoalCompletion:        goalCompletion,
		GoalCompletionTrend:   "0%",
	}, nil
}
