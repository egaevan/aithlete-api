package schedule

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	createScheduleUseCase      usecase.CreateScheduleUseCase
	getScheduleUseCase         usecase.GetScheduleUseCase
	listSchedulesUseCase       usecase.ListSchedulesUseCase
	listSchedulesByDateUseCase usecase.ListSchedulesByDateUseCase
	updateScheduleUseCase      usecase.UpdateScheduleUseCase
	deleteScheduleUseCase      usecase.DeleteScheduleUseCase
	toggleScheduleUseCase      usecase.ToggleScheduleUseCase
}

func New(
	create usecase.CreateScheduleUseCase,
	get usecase.GetScheduleUseCase,
	list usecase.ListSchedulesUseCase,
	listByDate usecase.ListSchedulesByDateUseCase,
	update usecase.UpdateScheduleUseCase,
	delete usecase.DeleteScheduleUseCase,
	toggle usecase.ToggleScheduleUseCase,
) *Handler {
	return &Handler{
		createScheduleUseCase:      create,
		getScheduleUseCase:         get,
		listSchedulesUseCase:       list,
		listSchedulesByDateUseCase: listByDate,
		updateScheduleUseCase:      update,
		deleteScheduleUseCase:      delete,
		toggleScheduleUseCase:      toggle,
	}
}
