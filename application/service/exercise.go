package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsExerciseNotFound(err error) bool {
	return errors.Is(err, domainerr.ErrExerciseNotFound)
}
