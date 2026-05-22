package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type WorkoutHandler struct {
	provider *mock.MockProvider
}

func NewWorkoutHandler(provider *mock.MockProvider) *WorkoutHandler {
	return &WorkoutHandler{provider: provider}
}

func (h *WorkoutHandler) GetWorkouts(c echo.Context) error {
	data := h.provider.GetWorkouts()
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *WorkoutHandler) GetWorkout(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.GetWorkout(id)
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *WorkoutHandler) CreateWorkout(c echo.Context) error {
	var req request.CreateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.CreateWorkout(req.Name, req.Date, req.WeightUnit, req.Notes)
	return response.Success(c, 201, code.Success, message.MsgWorkoutCreated, data)
}

func (h *WorkoutHandler) UpdateWorkout(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.UpdateWorkout(id, req.Name, req.Date)
	return response.Success(c, 200, code.Success, message.MsgWorkoutUpdated, data)
}

func (h *WorkoutHandler) DeleteWorkout(c echo.Context) error {
	data := h.provider.DeleteWorkout()
	return response.Success(c, 200, code.Success, message.MsgWorkoutDeleted, data)
}

func (h *WorkoutHandler) GetWorkoutStats(c echo.Context) error {
	data := h.provider.GetWorkoutStats()
	return response.SuccessOK(c, 200, code.Success, data)
}
