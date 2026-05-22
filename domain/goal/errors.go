package goal

import "errors"

var (
	ErrGoalNotFound          = errors.New("goal not found")
	ErrGoalAlreadyCompleted  = errors.New("goal already completed")
	ErrInvalidGoalTarget     = errors.New("goal target must be positive")
)
