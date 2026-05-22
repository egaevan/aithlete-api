package goal

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	createGoalUseCase         usecase.CreateGoalUseCase
	getGoalUseCase            usecase.GetGoalUseCase
	listGoalsUseCase          usecase.ListGoalsUseCase
	updateGoalUseCase         usecase.UpdateGoalUseCase
	deleteGoalUseCase         usecase.DeleteGoalUseCase
	toggleGoalUseCase         usecase.ToggleGoalUseCase
	updateGoalProgressUseCase usecase.UpdateGoalProgressUseCase
}

func New(
	create usecase.CreateGoalUseCase,
	get usecase.GetGoalUseCase,
	list usecase.ListGoalsUseCase,
	update usecase.UpdateGoalUseCase,
	delete usecase.DeleteGoalUseCase,
	toggle usecase.ToggleGoalUseCase,
	updateProgress usecase.UpdateGoalProgressUseCase,
) *Handler {
	return &Handler{
		createGoalUseCase:         create,
		getGoalUseCase:            get,
		listGoalsUseCase:          list,
		updateGoalUseCase:         update,
		deleteGoalUseCase:         delete,
		toggleGoalUseCase:         toggle,
		updateGoalProgressUseCase: updateProgress,
	}
}
