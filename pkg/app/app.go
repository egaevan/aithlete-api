package app

import (
	"context"

	"github.com/aithlete/aithlete-api/application/usecase"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/infrastructure/auth"
	"github.com/aithlete/aithlete-api/infrastructure/config"
	"github.com/aithlete/aithlete-api/infrastructure/database"
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	authhandler "github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	"github.com/aithlete/aithlete-api/interfaces/router"
	"github.com/aithlete/aithlete-api/pkg/mock"
)

type Dependencies struct {
	Config   *config.Config
	Handlers router.Handlers
}

func Bootstrap(log *logger.Logger) Dependencies {
	cfg := config.Load()

	hashSvc := auth.NewPasswordHasher()
	tokenSvc := auth.NewTokenService(cfg.Auth.JWTSecret, cfg.Auth.JWTExpiration)

	pool, err := database.NewPool(context.Background(), database.DSN(cfg.Database))
	if err != nil {
		log.Warn("Database not available, falling back to in-memory store: %v", err)
	}

	var userRepo repository.UserRepository
	if pool != nil {
		userRepo = database.NewUserRepository(pool)
	} else {
		userRepo = database.NewInMemoryUserRepository()
	}

	registerUC := usecase.NewRegisterUseCase(userRepo, hashSvc, tokenSvc)
	loginUC := usecase.NewLoginUseCase(userRepo, hashSvc, tokenSvc)
	refreshUC := usecase.NewRefreshTokenUseCase(tokenSvc)
	getMeUC := usecase.NewGetMeUseCase(userRepo)

	provider := mock.NewMockProvider()

	return Dependencies{
		Config: cfg,
		Handlers: router.Handlers{
			Auth:      authhandler.New(loginUC, registerUC, refreshUC, getMeUC),
			Workout:   handler.NewWorkoutHandler(provider),
			Exercise:  handler.NewExerciseHandler(provider),
			Progress:  handler.NewProgressHandler(provider),
			AI:        handler.NewAIHandler(provider),
			Analytics: handler.NewAnalyticsHandler(provider),
			Schedule:  handler.NewScheduleHandler(provider),
			Goal:      handler.NewGoalHandler(provider),
			Profile:   handler.NewProfileHandler(provider),
		},
	}
}
