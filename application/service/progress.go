package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsNoProgressData(err error) bool {
	return errors.Is(err, domainerr.ErrNoProgressData)
}
