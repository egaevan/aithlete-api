package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return response.Success(c, 200, "00000", "OK", map[string]any{
		"status":  "healthy",
		"service": "aithlete-api",
		"version": "0.1.0",
	})
}
