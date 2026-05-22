package exercise

import "github.com/aithlete/aithlete-api/application/usecase/exercise"

type Handler struct {
	listExercisesUseCase   exercise.ListExercisesUseCase
	getExerciseUseCase     exercise.GetExerciseUseCase
	listMuscleGroupsUseCase exercise.ListMuscleGroupsUseCase
}

func New(
	list exercise.ListExercisesUseCase,
	get exercise.GetExerciseUseCase,
	muscleGroups exercise.ListMuscleGroupsUseCase,
) *Handler {
	return &Handler{
		listExercisesUseCase:   list,
		getExerciseUseCase:     get,
		listMuscleGroupsUseCase: muscleGroups,
	}
}
