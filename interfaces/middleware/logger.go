package middleware

import (
	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

func RequestLogger(log *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Info("%s %s", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	}
}
