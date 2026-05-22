package router

import (
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	"github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	"github.com/aithlete/aithlete-api/interfaces/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Auth      *auth.Handler
	Workout   *handler.WorkoutHandler
	Exercise  *handler.ExerciseHandler
	Progress  *handler.ProgressHandler
	AI        *handler.AIHandler
	Analytics *handler.AnalyticsHandler
	Schedule  *handler.ScheduleHandler
	Goal      *handler.GoalHandler
	Profile   *handler.ProfileHandler
}

func New(log *logger.Logger, h Handlers) *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.CORS())
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.RequestLogger(log))

	healthHandler := handler.NewHealthHandler()

	e.GET("/health", healthHandler.HealthCheck)

	v1 := e.Group("/api/v1")
	{
		registerAuthRoutes(v1, h.Auth)
		registerWorkoutRoutes(v1, h.Workout)
		registerExerciseRoutes(v1, h.Exercise)
		registerProgressRoutes(v1, h.Progress)
		registerAIRoutes(v1, h.AI)
		registerAnalyticsRoutes(v1, h.Analytics)
		registerScheduleRoutes(v1, h.Schedule)
		registerGoalRoutes(v1, h.Goal)
		registerProfileRoutes(v1, h.Profile)
	}

	return e
}
