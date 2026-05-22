package service

import (
	"errors"

	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func IsNoAnalyticsData(err error) bool {
	return errors.Is(err, domainerr.ErrNoAnalyticsData)
}
