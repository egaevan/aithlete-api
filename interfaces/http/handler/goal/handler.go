package goal

import "github.com/aithlete/aithlete-api/application/usecase/goal"

type Handler struct {
	createGoalUseCase         goal.CreateGoalUseCase
	getGoalUseCase            goal.GetGoalUseCase
	listGoalsUseCase          goal.ListGoalsUseCase
	updateGoalUseCase         goal.UpdateGoalUseCase
	deleteGoalUseCase         goal.DeleteGoalUseCase
	toggleGoalUseCase         goal.ToggleGoalUseCase
	updateGoalProgressUseCase goal.UpdateGoalProgressUseCase
}

func New(
	create goal.CreateGoalUseCase,
	get goal.GetGoalUseCase,
	list goal.ListGoalsUseCase,
	update goal.UpdateGoalUseCase,
	delete goal.DeleteGoalUseCase,
	toggle goal.ToggleGoalUseCase,
	updateProgress goal.UpdateGoalProgressUseCase,
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
