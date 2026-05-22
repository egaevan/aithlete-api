package app

import (
	"context"

	"github.com/aithlete/aithlete-api/application/usecase"
	"github.com/aithlete/aithlete-api/domain/entity"
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

func Bootstrap(l *logger.Logger) Dependencies {
	cfg := config.Load()

	hashSvc := auth.NewPasswordHasher()
	tokenSvc := auth.NewTokenService(cfg.Auth.JWTSecret, cfg.Auth.JWTExpiration)

	pool, err := database.NewPool(context.Background(), database.DSN(cfg.Database))
	if err != nil {
		l.Warn("Database unavailable, using mock repo: %v", err)
	}

	var userRepo repository.UserRepository
	if pool != nil {
		dbRepo := database.NewUserRepository(pool)
		seedDefaultUser(dbRepo, hashSvc, l)
		userRepo = dbRepo
	} else {
		userRepo = repository.NewMockUserRepository()
	}

	registerUC := usecase.NewRegisterUseCase(userRepo, hashSvc, tokenSvc)
	loginUC := usecase.NewLoginUseCase(userRepo, hashSvc, tokenSvc)
	refreshUC := usecase.NewRefreshTokenUseCase(tokenSvc)
	getMeUC := usecase.NewGetMeUseCase(userRepo)

	provider := mock.NewMockProvider()

	return Dependencies{
		Config: cfg,
		Handlers: router.Handlers{
			TokenSvc: tokenSvc,
			Auth:     authhandler.New(loginUC, registerUC, refreshUC, getMeUC),
			Workout:  handler.NewWorkoutHandler(provider),
			Exercise: handler.NewExerciseHandler(provider),
			Progress: handler.NewProgressHandler(provider),
			AI:       handler.NewAIHandler(provider),
			Analytics: handler.NewAnalyticsHandler(provider),
			Schedule: handler.NewScheduleHandler(provider),
			Goal:     handler.NewGoalHandler(provider),
			Profile:  handler.NewProfileHandler(provider),
		},
	}
}

func seedDefaultUser(repo repository.UserRepository, hashSvc *auth.PasswordHasher, l *logger.Logger) {
	const defaultEmail = "alex@example.com"

	_, err := repo.FindByEmail(context.Background(), defaultEmail)
	if err == nil {
		return
	}

	hashed, err := hashSvc.Hash(defaultEmail)
	if err != nil {
		l.Warn("Failed to hash default password: %v", err)
		return
	}

	u := entity.NewUser(defaultEmail, "Alex Johnson", hashed)
	if err := repo.Create(context.Background(), u); err != nil {
		l.Warn("Failed to seed default user: %v", err)
		return
	}

	l.Info("Seeded default user: alex@example.com / alex@example.com")
}
