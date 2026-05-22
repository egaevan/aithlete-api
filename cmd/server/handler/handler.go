package handler

import (
	"net/http"

	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/infrastructure/auth"
	"github.com/aithlete/aithlete-api/infrastructure/config"
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/application/usecase"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	authhandler "github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	"github.com/aithlete/aithlete-api/interfaces/router"
	"github.com/aithlete/aithlete-api/pkg/mock"
)

var h http.Handler

func init() {
	log := logger.New()
	cfg := config.Load()

	hashSvc := auth.NewPasswordHasher()
	tokenSvc := auth.NewTokenService(cfg.Auth.JWTSecret, cfg.Auth.JWTExpiration)
	userRepo := repository.NewMockUserRepository()

	registerUC := usecase.NewRegisterUseCase(userRepo, hashSvc, tokenSvc)
	loginUC := usecase.NewLoginUseCase(userRepo, hashSvc, tokenSvc)
	refreshUC := usecase.NewRefreshTokenUseCase(tokenSvc)
	getMeUC := usecase.NewGetMeUseCase(userRepo)

	provider := mock.NewMockProvider()

	handlers := router.Handlers{
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
	}

	h = router.New(log, handlers)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}
