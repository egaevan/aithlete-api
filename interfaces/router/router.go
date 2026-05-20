package router

import (
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	"github.com/aithlete/aithlete-api/interfaces/middleware"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func New(log *logger.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.CORS())
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.RequestLogger(log))

	provider := mock.NewMockProvider()

	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(provider)
	workoutHandler := handler.NewWorkoutHandler(provider)
	exerciseHandler := handler.NewExerciseHandler(provider)
	progressHandler := handler.NewProgressHandler(provider)
	aiHandler := handler.NewAIHandler(provider)
	analyticsHandler := handler.NewAnalyticsHandler(provider)
	scheduleHandler := handler.NewScheduleHandler(provider)
	goalHandler := handler.NewGoalHandler(provider)
	profileHandler := handler.NewProfileHandler(provider)

	e.GET("/health", healthHandler.HealthCheck)

	v1 := e.Group("/api/v1")
	{
		registerAuthRoutes(v1, authHandler)
		registerWorkoutRoutes(v1, workoutHandler)
		registerExerciseRoutes(v1, exerciseHandler)
		registerProgressRoutes(v1, progressHandler)
		registerAIRoutes(v1, aiHandler)
		registerAnalyticsRoutes(v1, analyticsHandler)
		registerScheduleRoutes(v1, scheduleHandler)
		registerGoalRoutes(v1, goalHandler)
		registerProfileRoutes(v1, profileHandler)
	}

	return e
}

func registerAuthRoutes(g *echo.Group, h *handler.AuthHandler) {
	auth := g.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
	auth.POST("/logout", h.Logout)
	auth.GET("/me", h.GetMe)
	auth.POST("/refresh", h.RefreshToken)
}

func registerWorkoutRoutes(g *echo.Group, h *handler.WorkoutHandler) {
	workouts := g.Group("/workouts")
	workouts.GET("", h.GetWorkouts)
	workouts.POST("", h.CreateWorkout)
	workouts.GET("/stats", h.GetWorkoutStats)
	workouts.GET("/:id", h.GetWorkout)
	workouts.PUT("/:id", h.UpdateWorkout)
	workouts.DELETE("/:id", h.DeleteWorkout)
}

func registerExerciseRoutes(g *echo.Group, h *handler.ExerciseHandler) {
	exercises := g.Group("/exercises")
	exercises.GET("", h.GetExercises)
	exercises.GET("/muscle-groups", h.GetMuscleGroups)
	exercises.GET("/:id", h.GetExercise)
}

func registerProgressRoutes(g *echo.Group, h *handler.ProgressHandler) {
	progress := g.Group("/progress")
	progress.GET("/body-weight", h.GetBodyWeightHistory)
	progress.POST("/body-weight", h.AddBodyWeight)
	progress.GET("/strength", h.GetStrengthProgression)
	progress.GET("/consistency", h.GetConsistency)
	progress.GET("/muscle-volume", h.GetMuscleVolume)
	progress.GET("/overview", h.GetProgressOverview)
}

func registerAIRoutes(g *echo.Group, h *handler.AIHandler) {
	ai := g.Group("/ai")
	ai.GET("/recommendations", h.GetRecommendations)
	ai.POST("/chat", h.CreateChatSession)
	ai.GET("/chat/:sessionId", h.GetChatHistory)
	ai.POST("/chat/:sessionId", h.SendChatMessage)
	ai.GET("/fatigue", h.GetFatigueAnalysis)
	ai.GET("/recovery", h.GetRecoveryScore)
	ai.GET("/plateau", h.GetPlateauDetection)
}

func registerAnalyticsRoutes(g *echo.Group, h *handler.AnalyticsHandler) {
	analytics := g.Group("/analytics")
	analytics.GET("/dashboard", h.GetDashboard)
	analytics.GET("/weekly", h.GetWeeklyProgress)
	analytics.GET("/streak", h.GetStreak)
	analytics.GET("/overview", h.GetOverview)
	analytics.GET("/volume/weekly", h.GetWeeklyVolume)
	analytics.GET("/volume/muscle", h.GetMuscleVolumeDistribution)
}

func registerScheduleRoutes(g *echo.Group, h *handler.ScheduleHandler) {
	schedules := g.Group("/schedules")
	schedules.GET("", h.GetSchedules)
	schedules.GET("/today", h.GetTodaySchedules)
	schedules.POST("", h.CreateSchedule)
	schedules.GET("/:id", h.GetSchedule)
	schedules.PUT("/:id", h.UpdateSchedule)
	schedules.DELETE("/:id", h.DeleteSchedule)
	schedules.PATCH("/:id/toggle", h.ToggleSchedule)
}

func registerGoalRoutes(g *echo.Group, h *handler.GoalHandler) {
	goals := g.Group("/goals")
	goals.GET("", h.GetGoals)
	goals.POST("", h.CreateGoal)
	goals.GET("/:id", h.GetGoal)
	goals.PUT("/:id", h.UpdateGoal)
	goals.DELETE("/:id", h.DeleteGoal)
	goals.PATCH("/:id/toggle", h.ToggleGoal)
	goals.PATCH("/:id/progress", h.UpdateGoalProgress)
}

func registerProfileRoutes(g *echo.Group, h *handler.ProfileHandler) {
	g.PUT("/profile", h.UpdateProfile)
}
