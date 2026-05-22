package app

import (
	"context"

	authuc "github.com/aithlete/aithlete-api/application/usecase/auth"
	exerciseuc "github.com/aithlete/aithlete-api/application/usecase/exercise"
	goaluc "github.com/aithlete/aithlete-api/application/usecase/goal"
	profileuc "github.com/aithlete/aithlete-api/application/usecase/profile"
	progressuc "github.com/aithlete/aithlete-api/application/usecase/progress"
	scheduleuc "github.com/aithlete/aithlete-api/application/usecase/schedule"
	workoutuc "github.com/aithlete/aithlete-api/application/usecase/workout"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/infrastructure/auth"
	"github.com/aithlete/aithlete-api/infrastructure/config"
	"github.com/aithlete/aithlete-api/infrastructure/database"
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/http/handler"
	authhandler "github.com/aithlete/aithlete-api/interfaces/http/handler/auth"
	exercisehandler "github.com/aithlete/aithlete-api/interfaces/http/handler/exercise"
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

	registerUC := authuc.NewRegisterUseCase(userRepo, hashSvc, tokenSvc)
	loginUC := authuc.NewLoginUseCase(userRepo, hashSvc, tokenSvc)
	refreshUC := authuc.NewRefreshTokenUseCase(tokenSvc)
	getMeUC := authuc.NewGetMeUseCase(userRepo)
	updateProfileUC := profileuc.NewUpdateProfileUseCase(userRepo)

	var workoutRepo repository.WorkoutRepository
	if pool != nil {
		workoutRepo = database.NewWorkoutRepository(pool)
	} else {
		workoutRepo = repository.NewMockWorkoutRepository()
	}

	createWorkoutUC := workoutuc.NewCreateWorkoutUseCase(workoutRepo)
	getWorkoutUC := workoutuc.NewGetWorkoutUseCase(workoutRepo)
	listWorkoutsUC := workoutuc.NewListWorkoutsUseCase(workoutRepo)
	updateWorkoutUC := workoutuc.NewUpdateWorkoutUseCase(workoutRepo)
	deleteWorkoutUC := workoutuc.NewDeleteWorkoutUseCase(workoutRepo)
	completeWorkoutUC := workoutuc.NewCompleteWorkoutUseCase(workoutRepo)
	addExerciseUC := workoutuc.NewAddExerciseUseCase(workoutRepo)
	updateSetUC := workoutuc.NewUpdateSetUseCase(workoutRepo)

	var scheduleRepo repository.ScheduleRepository
	if pool != nil {
		scheduleRepo = database.NewScheduleRepository(pool)
	} else {
		scheduleRepo = repository.NewMockScheduleRepository()
	}

	createScheduleUC := scheduleuc.NewCreateScheduleUseCase(scheduleRepo)
	getScheduleUC := scheduleuc.NewGetScheduleUseCase(scheduleRepo)
	listSchedulesUC := scheduleuc.NewListSchedulesUseCase(scheduleRepo)
	listSchedulesByDateUC := scheduleuc.NewListSchedulesByDateUseCase(scheduleRepo)
	updateScheduleUC := scheduleuc.NewUpdateScheduleUseCase(scheduleRepo)
	deleteScheduleUC := scheduleuc.NewDeleteScheduleUseCase(scheduleRepo)
	toggleScheduleUC := scheduleuc.NewToggleScheduleUseCase(scheduleRepo)

	var progressRepo repository.ProgressRepository
	if pool != nil {
		progressRepo = database.NewProgressRepository(pool)
	} else {
		progressRepo = repository.NewMockProgressRepository()
	}

	getBodyWeightHistoryUC := progressuc.NewGetBodyWeightHistoryUseCase(progressRepo)
	addBodyWeightUC := progressuc.NewAddBodyWeightUseCase(progressRepo)
	getStrengthProgressionUC := progressuc.NewGetStrengthProgressionUseCase(progressRepo)
	getConsistencyUC := progressuc.NewGetConsistencyUseCase(progressRepo)
	getMuscleVolumeUC := progressuc.NewGetMuscleVolumeUseCase(progressRepo)
	getProgressOverviewUC := progressuc.NewGetProgressOverviewUseCase(progressRepo)

	var goalRepo repository.GoalRepository
	if pool != nil {
		goalRepo = database.NewGoalRepository(pool)
	} else {
		goalRepo = repository.NewMockGoalRepository()
	}

	createGoalUC := goaluc.NewCreateGoalUseCase(goalRepo)
	getGoalUC := goaluc.NewGetGoalUseCase(goalRepo)
	listGoalsUC := goaluc.NewListGoalsUseCase(goalRepo)
	updateGoalUC := goaluc.NewUpdateGoalUseCase(goalRepo)
	deleteGoalUC := goaluc.NewDeleteGoalUseCase(goalRepo)
	toggleGoalUC := goaluc.NewToggleGoalUseCase(goalRepo)
	updateGoalProgressUC := goaluc.NewUpdateGoalProgressUseCase(goalRepo)

	var exerciseRepo repository.ExerciseRepository
	if pool != nil {
		exerciseRepo = database.NewExerciseRepository(pool)
	} else {
		exerciseRepo = repository.NewMockExerciseRepository()
	}

	listExercisesUC := exerciseuc.NewListExercisesUseCase(exerciseRepo)
	getExerciseUC := exerciseuc.NewGetExerciseUseCase(exerciseRepo)
	listMuscleGroupsUC := exerciseuc.NewListMuscleGroupsUseCase()

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
			Exercise: exercisehandler.New(listExercisesUC, getExerciseUC, listMuscleGroupsUC),
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
