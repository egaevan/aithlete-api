package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsWorkoutNotFound(err error) bool {
	return errors.Is(err, domainerr.ErrWorkoutNotFound)
}
