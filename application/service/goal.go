package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsGoalNotFound(err error) bool {
	return errors.Is(err, domainerr.ErrGoalNotFound)
}
