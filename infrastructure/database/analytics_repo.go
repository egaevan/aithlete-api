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

	var (
		caloriesBurned int
		activeMinutes  int
	)
	err := r.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(c.duration * 10), 0), COALESCE(SUM(c.duration), 0)
		FROM completions c
		JOIN workouts w ON w.id = c.workout_id
		WHERE w.user_id = $1 AND c.completed_at >= $2
	`, userID, weekStart).Scan(&caloriesBurned, &activeMinutes)
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
		RecentActivity: nil,
		MuscleRecovery: nil,
		TodaySchedule:  nil,
	}, nil
}

func (r *AnalyticsRepository) GetWeeklyProgress(ctx context.Context, userID string) ([]entity.WeeklyProgressDay, error) {
	return r.getWeeklyProgress(ctx, userID, time.Now())
}

func (r *AnalyticsRepository) getWeeklyProgress(ctx context.Context, userID string, now time.Time) ([]entity.WeeklyProgressDay, error) {
	weekStart := now.AddDate(0, 0, -int(now.Weekday()))

	rows, err := r.pool.Query(ctx, `
		SELECT
			to_char(c.completed_at, 'Day') AS day,
			COALESCE(SUM(c.duration * 10), 0) AS calories,
			COALESCE(SUM(c.duration), 0) AS duration
		FROM completions c
		JOIN workouts w ON w.id = c.workout_id
		WHERE w.user_id = $1 AND c.completed_at >= $2
		GROUP BY to_char(c.completed_at, 'Day'), EXTRACT(DOW FROM c.completed_at)
		ORDER BY EXTRACT(DOW FROM c.completed_at)
	`, userID, weekStart)
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
		WITH daily AS (
			SELECT DISTINCT DATE(c.completed_at) AS work_date
			FROM completions c
			JOIN workouts w ON w.id = c.workout_id
			WHERE w.user_id = $1
		),
		streaks AS (
			SELECT work_date,
				work_date - ROW_NUMBER() OVER (ORDER BY work_date)::int AS grp
			FROM daily
		)
		SELECT COUNT(*) AS current_streak
		FROM streaks
		GROUP BY grp
		ORDER BY MAX(work_date) DESC
		LIMIT 1
	`, userID).Scan(&current)
	if err != nil {
		return nil, domainerr.ErrNoAnalyticsData
	}

	return &entity.Streak{
		Current:      current,
		PersonalBest: current,
		History:      nil,
	}, nil
}

func (r *AnalyticsRepository) GetWeeklyVolume(ctx context.Context, userID string) ([]entity.WeeklyVolume, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			to_char(c.completed_at, 'IYYY"-W"IW') AS week,
			COALESCE(SUM(c.volume), 0) AS volume
		FROM completions c
		JOIN workouts w ON w.id = c.workout_id
		WHERE w.user_id = $1
		GROUP BY to_char(c.completed_at, 'IYYY"-W"IW')
		ORDER BY to_char(c.completed_at, 'IYYY"-W"IW') DESC
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
		SELECT
			e.muscle_group,
			COALESCE(SUM(ce.volume), 0) AS volume
		FROM completions c
		JOIN workouts w ON w.id = c.workout_id
		JOIN completion_exercises ce ON ce.completion_id = c.id
		JOIN exercises e ON e.id = ce.exercise_id
		WHERE w.user_id = $1
		GROUP BY e.muscle_group
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
	var (
		totalVolume       int
		avgSession        int
		sessionsPerMonth  int
		goalCompletion    int
	)
	err := r.pool.QueryRow(ctx, `
		SELECT
			COALESCE(SUM(c.volume), 0),
			COALESCE(ROUND(AVG(c.duration)), 0),
			COUNT(DISTINCT c.workout_id)
		FROM completions c
		JOIN workouts w ON w.id = c.workout_id
		WHERE w.user_id = $1
		  AND c.completed_at >= NOW() - INTERVAL '30 days'
	`, userID).Scan(&totalVolume, &avgSession, &sessionsPerMonth)
	if err != nil {
		return nil, fmt.Errorf("query overview: %w", err)
	}

	err = r.pool.QueryRow(ctx, `
		SELECT COALESCE(ROUND(AVG(CASE WHEN completed THEN 100.0 ELSE 0 END)), 0)
		FROM goals WHERE user_id = $1
	`, userID).Scan(&goalCompletion)
	if err != nil {
		return nil, fmt.Errorf("query goal completion: %w", err)
	}

	return &entity.AnalyticsOverview{
		TotalVolume:           totalVolume,
		TotalVolumeTrend:     "0%",
		AvgSession:            avgSession,
		AvgSessionTrend:      "0 min",
		SessionsPerMonth:      sessionsPerMonth,
		SessionsPerMonthTrend: "0",
		GoalCompletion:        goalCompletion,
		GoalCompletionTrend:  "0%",
	}, nil
}
