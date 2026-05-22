package workout

import "errors"

var (
	ErrWorkoutNotFound    = errors.New("workout not found")
	ErrEmptyWorkout       = errors.New("cannot complete empty workout")
	ErrDuplicateExercise  = errors.New("exercise already exists in workout")
	ErrExerciseNotFound   = errors.New("exercise not found in workout")
	ErrSetNotFound        = errors.New("set not found")
	ErrInvalidSetValue    = errors.New("reps and weight must be positive")
	ErrIncompleteWorkout  = errors.New("all sets must be completed")
)
