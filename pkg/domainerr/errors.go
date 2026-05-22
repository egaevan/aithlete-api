package domainerr

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

var (
	ErrWorkoutNotFound          = errors.New("workout not found")
	ErrEmptyWorkout             = errors.New("cannot complete empty workout")
	ErrDuplicateExercise        = errors.New("exercise already exists in workout")
	ErrExerciseNotFoundInWorkout = errors.New("exercise not found in workout")
	ErrSetNotFound              = errors.New("set not found")
	ErrInvalidSetValue          = errors.New("reps and weight must be positive")
	ErrIncompleteWorkout        = errors.New("all sets must be completed")
)

var ErrExerciseNotFound = errors.New("exercise not found")

var (
	ErrScheduleNotFound = errors.New("schedule not found")
	ErrScheduleConflict = errors.New("schedule time conflict")
)

var (
	ErrGoalNotFound         = errors.New("goal not found")
	ErrGoalAlreadyCompleted = errors.New("goal already completed")
	ErrInvalidGoalTarget    = errors.New("goal target must be positive")
)

var ErrNoProgressData = errors.New("no progress data available")

var ErrSessionNotFound = errors.New("chat session not found")

var ErrNoAnalyticsData = errors.New("no analytics data available")
