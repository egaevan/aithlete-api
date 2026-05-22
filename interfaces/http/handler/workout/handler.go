package workout

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	createWorkoutUseCase    usecase.CreateWorkoutUseCase
	getWorkoutUseCase       usecase.GetWorkoutUseCase
	listWorkoutsUseCase     usecase.ListWorkoutsUseCase
	updateWorkoutUseCase    usecase.UpdateWorkoutUseCase
	deleteWorkoutUseCase    usecase.DeleteWorkoutUseCase
	completeWorkoutUseCase  usecase.CompleteWorkoutUseCase
	addExerciseUseCase      usecase.AddExerciseUseCase
	updateSetUseCase        usecase.UpdateSetUseCase
}

func New(
	create usecase.CreateWorkoutUseCase,
	get usecase.GetWorkoutUseCase,
	list usecase.ListWorkoutsUseCase,
	update usecase.UpdateWorkoutUseCase,
	delete usecase.DeleteWorkoutUseCase,
	complete usecase.CompleteWorkoutUseCase,
	addExercise usecase.AddExerciseUseCase,
	updateSet usecase.UpdateSetUseCase,
) *Handler {
	return &Handler{
		createWorkoutUseCase:   create,
		getWorkoutUseCase:      get,
		listWorkoutsUseCase:    list,
		updateWorkoutUseCase:   update,
		deleteWorkoutUseCase:   delete,
		completeWorkoutUseCase: complete,
		addExerciseUseCase:     addExercise,
		updateSetUseCase:       updateSet,
	}
}
