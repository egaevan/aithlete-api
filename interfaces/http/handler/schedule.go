package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	provider *mock.MockProvider
}

func NewScheduleHandler(provider *mock.MockProvider) *ScheduleHandler {
	return &ScheduleHandler{provider: provider}
}

func (h *ScheduleHandler) GetSchedules(c echo.Context) error {
	data := h.provider.GetSchedules()
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *ScheduleHandler) GetTodaySchedules(c echo.Context) error {
	data := h.provider.GetTodaySchedules()
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *ScheduleHandler) GetSchedule(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.GetSchedule(id)
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *ScheduleHandler) CreateSchedule(c echo.Context) error {
	var req request.CreateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.CreateSchedule(req.Date, req.Time, req.Title, req.Duration, req.Type, req.Notes)
	return response.Success(c, 201, "00000", "Created", data)
}

func (h *ScheduleHandler) UpdateSchedule(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.UpdateSchedule(id, req.Date, req.Time, req.Title, req.Duration, req.Type, req.Notes)
	return response.Success(c, 200, "00000", "OK", data)
}

func (h *ScheduleHandler) DeleteSchedule(c echo.Context) error {
	data := h.provider.DeleteSchedule()
	return response.Success(c, 200, "00000", "Deleted", data)
}

func (h *ScheduleHandler) ToggleSchedule(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.ToggleSchedule(id)
	return response.Success(c, 200, "00000", "OK", data)
}
