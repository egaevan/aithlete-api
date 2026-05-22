package router

import (
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	"github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	"github.com/labstack/echo/v4"
)

func registerAuthRoutes(g *echo.Group, h *auth.Handler, authMw echo.MiddlewareFunc) {
	auth := g.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
	auth.POST("/refresh", h.RefreshToken)
	auth.POST("/logout", h.Logout, authMw)
	auth.GET("/me", h.GetMe, authMw)
}

func registerWorkoutRoutes(g *echo.Group, h *handler.WorkoutHandler, authMw echo.MiddlewareFunc) {
	workouts := g.Group("/workouts", authMw)
	workouts.GET("", h.GetWorkouts)
	workouts.POST("", h.CreateWorkout)
	workouts.GET("/stats", h.GetWorkoutStats)
	workouts.GET("/:id", h.GetWorkout)
	workouts.PUT("/:id", h.UpdateWorkout)
	workouts.DELETE("/:id", h.DeleteWorkout)
}

func registerExerciseRoutes(g *echo.Group, h *handler.ExerciseHandler, authMw echo.MiddlewareFunc) {
	exercises := g.Group("/exercises", authMw)
	exercises.GET("", h.GetExercises)
	exercises.GET("/muscle-groups", h.GetMuscleGroups)
	exercises.GET("/:id", h.GetExercise)
}

func registerProgressRoutes(g *echo.Group, h *handler.ProgressHandler, authMw echo.MiddlewareFunc) {
	progress := g.Group("/progress", authMw)
	progress.GET("/body-weight", h.GetBodyWeightHistory)
	progress.POST("/body-weight", h.AddBodyWeight)
	progress.GET("/strength", h.GetStrengthProgression)
	progress.GET("/consistency", h.GetConsistency)
	progress.GET("/muscle-volume", h.GetMuscleVolume)
	progress.GET("/overview", h.GetProgressOverview)
}

func registerAIRoutes(g *echo.Group, h *handler.AIHandler, authMw echo.MiddlewareFunc) {
	ai := g.Group("/ai", authMw)
	ai.GET("/recommendations", h.GetRecommendations)
	ai.POST("/chat", h.CreateChatSession)
	ai.GET("/chat/:sessionId", h.GetChatHistory)
	ai.POST("/chat/:sessionId", h.SendChatMessage)
	ai.GET("/fatigue", h.GetFatigueAnalysis)
	ai.GET("/recovery", h.GetRecoveryScore)
	ai.GET("/plateau", h.GetPlateauDetection)
}

func registerAnalyticsRoutes(g *echo.Group, h *handler.AnalyticsHandler, authMw echo.MiddlewareFunc) {
	analytics := g.Group("/analytics", authMw)
	analytics.GET("/dashboard", h.GetDashboard)
	analytics.GET("/weekly", h.GetWeeklyProgress)
	analytics.GET("/streak", h.GetStreak)
	analytics.GET("/overview", h.GetOverview)
	analytics.GET("/volume/weekly", h.GetWeeklyVolume)
	analytics.GET("/volume/muscle", h.GetMuscleVolumeDistribution)
}

func registerScheduleRoutes(g *echo.Group, h *handler.ScheduleHandler, authMw echo.MiddlewareFunc) {
	schedules := g.Group("/schedules", authMw)
	schedules.GET("", h.GetSchedules)
	schedules.GET("/today", h.GetTodaySchedules)
	schedules.POST("", h.CreateSchedule)
	schedules.GET("/:id", h.GetSchedule)
	schedules.PUT("/:id", h.UpdateSchedule)
	schedules.DELETE("/:id", h.DeleteSchedule)
	schedules.PATCH("/:id/toggle", h.ToggleSchedule)
}

func registerGoalRoutes(g *echo.Group, h *handler.GoalHandler, authMw echo.MiddlewareFunc) {
	goals := g.Group("/goals", authMw)
	goals.GET("", h.GetGoals)
	goals.POST("", h.CreateGoal)
	goals.GET("/:id", h.GetGoal)
	goals.PUT("/:id", h.UpdateGoal)
	goals.DELETE("/:id", h.DeleteGoal)
	goals.PATCH("/:id/toggle", h.ToggleGoal)
	goals.PATCH("/:id/progress", h.UpdateGoalProgress)
}

func registerProfileRoutes(g *echo.Group, h *handler.ProfileHandler, authMw echo.MiddlewareFunc) {
	g.PUT("/profile", h.UpdateProfile, authMw)
}
