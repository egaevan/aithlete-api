package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsScheduleNotFound(err error) bool {
	return errors.Is(err, domainerr.ErrScheduleNotFound)
}
