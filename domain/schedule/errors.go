package schedule

import "errors"

var (
	ErrScheduleNotFound = errors.New("schedule not found")
	ErrScheduleConflict = errors.New("schedule time conflict")
)
