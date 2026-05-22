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
	goalhandler "github.com/aithlete/aithlete-api/interfaces/http/handler/goal"
	progresshandler "github.com/aithlete/aithlete-api/interfaces/http/handler/progress"
	profilehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/profile"
	schedulehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/schedule"
	workouthandler "github.com/aithlete/aithlete-api/interfaces/http/handler/workout"
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
	updateProfileUC := usecase.NewUpdateProfileUseCase(userRepo)

	var workoutRepo repository.WorkoutRepository
	if pool != nil {
		workoutRepo = database.NewWorkoutRepository(pool)
	} else {
		workoutRepo = repository.NewMockWorkoutRepository()
	}

	createWorkoutUC := usecase.NewCreateWorkoutUseCase(workoutRepo)
	getWorkoutUC := usecase.NewGetWorkoutUseCase(workoutRepo)
	listWorkoutsUC := usecase.NewListWorkoutsUseCase(workoutRepo)
	updateWorkoutUC := usecase.NewUpdateWorkoutUseCase(workoutRepo)
	deleteWorkoutUC := usecase.NewDeleteWorkoutUseCase(workoutRepo)
	completeWorkoutUC := usecase.NewCompleteWorkoutUseCase(workoutRepo)
	addExerciseUC := usecase.NewAddExerciseUseCase(workoutRepo)
	updateSetUC := usecase.NewUpdateSetUseCase(workoutRepo)

	var scheduleRepo repository.ScheduleRepository
	if pool != nil {
		scheduleRepo = database.NewScheduleRepository(pool)
	} else {
		scheduleRepo = repository.NewMockScheduleRepository()
	}

	createScheduleUC := usecase.NewCreateScheduleUseCase(scheduleRepo)
	getScheduleUC := usecase.NewGetScheduleUseCase(scheduleRepo)
	listSchedulesUC := usecase.NewListSchedulesUseCase(scheduleRepo)
	listSchedulesByDateUC := usecase.NewListSchedulesByDateUseCase(scheduleRepo)
	updateScheduleUC := usecase.NewUpdateScheduleUseCase(scheduleRepo)
	deleteScheduleUC := usecase.NewDeleteScheduleUseCase(scheduleRepo)
	toggleScheduleUC := usecase.NewToggleScheduleUseCase(scheduleRepo)

	var progressRepo repository.ProgressRepository
	if pool != nil {
		progressRepo = database.NewProgressRepository(pool)
	} else {
		progressRepo = repository.NewMockProgressRepository()
	}

	getBodyWeightHistoryUC := usecase.NewGetBodyWeightHistoryUseCase(progressRepo)
	addBodyWeightUC := usecase.NewAddBodyWeightUseCase(progressRepo)
	getStrengthProgressionUC := usecase.NewGetStrengthProgressionUseCase(progressRepo)
	getConsistencyUC := usecase.NewGetConsistencyUseCase(progressRepo)
	getMuscleVolumeUC := usecase.NewGetMuscleVolumeUseCase(progressRepo)
	getProgressOverviewUC := usecase.NewGetProgressOverviewUseCase(progressRepo)

	var goalRepo repository.GoalRepository
	if pool != nil {
		goalRepo = database.NewGoalRepository(pool)
	} else {
		goalRepo = repository.NewMockGoalRepository()
	}

	createGoalUC := usecase.NewCreateGoalUseCase(goalRepo)
	getGoalUC := usecase.NewGetGoalUseCase(goalRepo)
	listGoalsUC := usecase.NewListGoalsUseCase(goalRepo)
	updateGoalUC := usecase.NewUpdateGoalUseCase(goalRepo)
	deleteGoalUC := usecase.NewDeleteGoalUseCase(goalRepo)
	toggleGoalUC := usecase.NewToggleGoalUseCase(goalRepo)
	updateGoalProgressUC := usecase.NewUpdateGoalProgressUseCase(goalRepo)

	provider := mock.NewMockProvider()

	return Dependencies{
		Config: cfg,
		Handlers: router.Handlers{
			TokenSvc: tokenSvc,
			Auth:     authhandler.New(loginUC, registerUC, refreshUC, getMeUC),
			Workout: workouthandler.New(
				createWorkoutUC, getWorkoutUC, listWorkoutsUC,
				updateWorkoutUC, deleteWorkoutUC, completeWorkoutUC,
				addExerciseUC, updateSetUC,
			),
			Profile:  profilehandler.New(updateProfileUC),
			Schedule: schedulehandler.New(
				createScheduleUC, getScheduleUC, listSchedulesUC,
				listSchedulesByDateUC, updateScheduleUC, deleteScheduleUC,
				toggleScheduleUC,
			),
			Progress: progresshandler.New(
				getBodyWeightHistoryUC, addBodyWeightUC, getStrengthProgressionUC,
				getConsistencyUC, getMuscleVolumeUC, getProgressOverviewUC,
			),
			Goal: goalhandler.New(
				createGoalUC, getGoalUC, listGoalsUC,
				updateGoalUC, deleteGoalUC, toggleGoalUC,
				updateGoalProgressUC,
			),
			Exercise: handler.NewExerciseHandler(provider),
			AI:       handler.NewAIHandler(provider),
			Analytics: handler.NewAnalyticsHandler(provider),
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
