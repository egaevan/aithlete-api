package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type AnalyticsHandler struct {
	provider *mock.MockProvider
}

func NewAnalyticsHandler(provider *mock.MockProvider) *AnalyticsHandler {
	return &AnalyticsHandler{provider: provider}
}

func (h *AnalyticsHandler) GetDashboard(c echo.Context) error {
	data := h.provider.GetDashboard()
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *AnalyticsHandler) GetWeeklyProgress(c echo.Context) error {
	data, meta := h.provider.GetWeeklyProgress()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, "00000", "Success", data, rMeta)
}

func (h *AnalyticsHandler) GetStreak(c echo.Context) error {
	data := h.provider.GetStreak()
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *AnalyticsHandler) GetOverview(c echo.Context) error {
	data := h.provider.GetAnalyticsOverview()
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *AnalyticsHandler) GetWeeklyVolume(c echo.Context) error {
	data := h.provider.GetWeeklyVolume()
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *AnalyticsHandler) GetMuscleVolumeDistribution(c echo.Context) error {
	data := h.provider.GetMuscleVolumeDistribution()
	return response.Success(c, 200, "00000", "OK", data)
}
