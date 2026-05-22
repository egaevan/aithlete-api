package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return response.SuccessOK(c, 200, code.Success, map[string]any{
		"status":  "healthy",
		"service": "aithlete-api",
		"version": "0.1.0",
	})
}
