package schedule

import "github.com/aithlete/aithlete-api/application/usecase/schedule"

type Handler struct {
	createScheduleUseCase      schedule.CreateScheduleUseCase
	getScheduleUseCase         schedule.GetScheduleUseCase
	listSchedulesUseCase       schedule.ListSchedulesUseCase
	listSchedulesByDateUseCase schedule.ListSchedulesByDateUseCase
	updateScheduleUseCase      schedule.UpdateScheduleUseCase
	deleteScheduleUseCase      schedule.DeleteScheduleUseCase
	toggleScheduleUseCase      schedule.ToggleScheduleUseCase
}

func New(
	create schedule.CreateScheduleUseCase,
	get schedule.GetScheduleUseCase,
	list schedule.ListSchedulesUseCase,
	listByDate schedule.ListSchedulesByDateUseCase,
	update schedule.UpdateScheduleUseCase,
	delete schedule.DeleteScheduleUseCase,
	toggle schedule.ToggleScheduleUseCase,
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
