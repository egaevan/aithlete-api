package router

import (
	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	analyticshandler "github.com/aithlete/aithlete-api/interfaces/http/handler/analytics"
	"github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	exercisehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/exercise"
	goalhandler "github.com/aithlete/aithlete-api/interfaces/http/handler/goal"
	progresshandler "github.com/aithlete/aithlete-api/interfaces/http/handler/progress"
	profilehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/profile"
	schedulehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/schedule"
	workouthandler "github.com/aithlete/aithlete-api/interfaces/http/handler/workout"
	"github.com/aithlete/aithlete-api/interfaces/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	TokenSvc service.TokenService
	Auth     *auth.Handler
	Workout  *workouthandler.Handler
	Profile  *profilehandler.Handler
	Progress *progresshandler.Handler
	Schedule *schedulehandler.Handler
	Goal     *goalhandler.Handler
	Exercise *exercisehandler.Handler
	AI       *handler.AIHandler
	Analytics *analyticshandler.Handler
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
	authMw := middleware.Auth(h.TokenSvc)
	{
		registerAuthRoutes(v1, h.Auth, authMw)
		registerWorkoutRoutes(v1, h.Workout, authMw)
		registerExerciseRoutes(v1, h.Exercise, authMw)
		registerProgressRoutes(v1, h.Progress, authMw)
		registerAIRoutes(v1, h.AI, authMw)
		registerAnalyticsRoutes(v1, h.Analytics, authMw)
		registerScheduleRoutes(v1, h.Schedule, authMw)
		registerGoalRoutes(v1, h.Goal, authMw)
		registerProfileRoutes(v1, h.Profile, authMw)
	}

	return e
}
