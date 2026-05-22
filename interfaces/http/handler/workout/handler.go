package workout

import "github.com/aithlete/aithlete-api/application/usecase/workout"

type Handler struct {
	createWorkoutUseCase    workout.CreateWorkoutUseCase
	getWorkoutUseCase       workout.GetWorkoutUseCase
	listWorkoutsUseCase     workout.ListWorkoutsUseCase
	updateWorkoutUseCase    workout.UpdateWorkoutUseCase
	deleteWorkoutUseCase    workout.DeleteWorkoutUseCase
	completeWorkoutUseCase  workout.CompleteWorkoutUseCase
	addExerciseUseCase      workout.AddExerciseUseCase
	updateSetUseCase        workout.UpdateSetUseCase
}

func New(
	create workout.CreateWorkoutUseCase,
	get workout.GetWorkoutUseCase,
	list workout.ListWorkoutsUseCase,
	update workout.UpdateWorkoutUseCase,
	delete workout.DeleteWorkoutUseCase,
	complete workout.CompleteWorkoutUseCase,
	addExercise workout.AddExerciseUseCase,
	updateSet workout.UpdateSetUseCase,
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
