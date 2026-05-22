package analytics

import "github.com/aithlete/aithlete-api/application/dto"

type DashboardResponse struct {
	Stats            DashboardStatsResponse            `json:"stats"`
	WeeklyProgress   []WeeklyProgressDayResponse       `json:"weeklyProgress"`
	Streak           StreakResponse                    `json:"streak"`
	RecentActivity   []ActivityResponse                `json:"recentActivity"`
	MuscleRecovery   []MuscleRecoveryResponse          `json:"muscleRecovery"`
	TodaySchedule    []TodayScheduleItemResponse       `json:"todaySchedule"`
	AIRecommendation *AIRecommendationSummaryResponse  `json:"aiRecommendation"`
}

type DashboardStatsResponse struct {
	CaloriesBurned     int    `json:"caloriesBurned"`
	CaloriesTrend      string `json:"caloriesTrend"`
	ActiveMinutes      int    `json:"activeMinutes"`
	ActiveMinutesTrend string `json:"activeMinutesTrend"`
	GoalsCompleted     int    `json:"goalsCompleted"`
	GoalsTotal         int    `json:"goalsTotal"`
	GoalsTrend         string `json:"goalsTrend"`
	AvgHeartRate       int    `json:"avgHeartRate"`
	HeartRateTrend     string `json:"heartRateTrend"`
}

type WeeklyProgressDayResponse struct {
	Day      string `json:"day"`
	Calories int    `json:"calories"`
	Duration int    `json:"duration"`
}

type StreakResponse struct {
	Current      int    `json:"current"`
	PersonalBest int    `json:"personalBest"`
	History      []bool `json:"history"`
}

type ActivityResponse struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

type MuscleRecoveryResponse struct {
	MuscleGroup      string `json:"muscleGroup"`
	Recovery         int    `json:"recovery"`
	ReadyForTraining bool   `json:"readyForTraining"`
}

type TodayScheduleItemResponse struct {
	ID        string `json:"id"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	Duration  string `json:"duration"`
	Type      string `json:"type"`
	Completed bool   `json:"completed"`
}

type AIRecommendationSummaryResponse struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Confidence  float64 `json:"confidence"`
}

type WeeklyVolumeResponse struct {
	Week   string  `json:"week"`
	Volume float64 `json:"volume"`
}

type MuscleVolumeDistributionResponse struct {
	Muscle string  `json:"muscle"`
	Volume float64 `json:"volume"`
}

type AnalyticsOverviewResponse struct {
	TotalVolume           int    `json:"totalVolume"`
	TotalVolumeTrend      string `json:"totalVolumeTrend"`
	AvgSession            int    `json:"avgSession"`
	AvgSessionTrend       string `json:"avgSessionTrend"`
	SessionsPerMonth      int    `json:"sessionsPerMonth"`
	SessionsPerMonthTrend string `json:"sessionsPerMonthTrend"`
	GoalCompletion        int    `json:"goalCompletion"`
	GoalCompletionTrend   string `json:"goalCompletionTrend"`
}

func toDashboardResponse(r *dto.DashboardResult) DashboardResponse {
	resp := DashboardResponse{
		Stats: DashboardStatsResponse{
			CaloriesBurned:     r.Stats.CaloriesBurned,
			CaloriesTrend:      r.Stats.CaloriesTrend,
			ActiveMinutes:      r.Stats.ActiveMinutes,
			ActiveMinutesTrend: r.Stats.ActiveMinutesTrend,
			GoalsCompleted:     r.Stats.GoalsCompleted,
			GoalsTotal:         r.Stats.GoalsTotal,
			GoalsTrend:         r.Stats.GoalsTrend,
			AvgHeartRate:       r.Stats.AvgHeartRate,
			HeartRateTrend:     r.Stats.HeartRateTrend,
		},
	}
	for _, w := range r.WeeklyProgress {
		resp.WeeklyProgress = append(resp.WeeklyProgress, WeeklyProgressDayResponse{
			Day: w.Day, Calories: w.Calories, Duration: w.Duration,
		})
	}
	resp.Streak = StreakResponse{
		Current: r.Streak.Current, PersonalBest: r.Streak.PersonalBest, History: r.Streak.History,
	}
	for _, a := range r.RecentActivity {
		resp.RecentActivity = append(resp.RecentActivity, ActivityResponse{
			ID: a.ID, Type: a.Type, Title: a.Title, Description: a.Description, Timestamp: a.Timestamp,
		})
	}
	for _, m := range r.MuscleRecovery {
		resp.MuscleRecovery = append(resp.MuscleRecovery, MuscleRecoveryResponse{
			MuscleGroup: m.MuscleGroup, Recovery: m.Recovery, ReadyForTraining: m.ReadyForTraining,
		})
	}
	for _, t := range r.TodaySchedule {
		resp.TodaySchedule = append(resp.TodaySchedule, TodayScheduleItemResponse{
			ID: t.ID, Time: t.Time, Title: t.Title, Duration: t.Duration, Type: t.Type, Completed: t.Completed,
		})
	}
	if r.AIRecommendation != nil {
		resp.AIRecommendation = &AIRecommendationSummaryResponse{
			ID: r.AIRecommendation.ID, Type: r.AIRecommendation.Type,
			Title: r.AIRecommendation.Title, Description: r.AIRecommendation.Description,
			Confidence: r.AIRecommendation.Confidence,
		}
	}
	return resp
}

func toWeeklyProgressDayResponseList(results []dto.WeeklyProgressDayResult) []WeeklyProgressDayResponse {
	resp := make([]WeeklyProgressDayResponse, len(results))
	for i, r := range results {
		resp[i] = WeeklyProgressDayResponse{Day: r.Day, Calories: r.Calories, Duration: r.Duration}
	}
	return resp
}

func toStreakResponse(r *dto.StreakResult) StreakResponse {
	return StreakResponse{Current: r.Current, PersonalBest: r.PersonalBest, History: r.History}
}

func toWeeklyVolumeResponseList(results []dto.WeeklyVolumeResult) []WeeklyVolumeResponse {
	resp := make([]WeeklyVolumeResponse, len(results))
	for i, r := range results {
		resp[i] = WeeklyVolumeResponse{Week: r.Week, Volume: r.Volume}
	}
	return resp
}

func toMuscleVolumeDistributionResponseList(results []dto.MuscleVolumeDistributionResult) []MuscleVolumeDistributionResponse {
	resp := make([]MuscleVolumeDistributionResponse, len(results))
	for i, r := range results {
		resp[i] = MuscleVolumeDistributionResponse{Muscle: r.Muscle, Volume: r.Volume}
	}
	return resp
}

func toAnalyticsOverviewResponse(r *dto.AnalyticsOverviewResult) AnalyticsOverviewResponse {
	return AnalyticsOverviewResponse{
		TotalVolume: r.TotalVolume, TotalVolumeTrend: r.TotalVolumeTrend,
		AvgSession: r.AvgSession, AvgSessionTrend: r.AvgSessionTrend,
		SessionsPerMonth: r.SessionsPerMonth, SessionsPerMonthTrend: r.SessionsPerMonthTrend,
		GoalCompletion: r.GoalCompletion, GoalCompletionTrend: r.GoalCompletionTrend,
	}
}
