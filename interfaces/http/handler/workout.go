package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
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
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *WorkoutHandler) GetWorkout(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.GetWorkout(id)
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *WorkoutHandler) CreateWorkout(c echo.Context) error {
	var req request.CreateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.CreateWorkout(req.Name, req.Date, req.WeightUnit, req.Notes)
	return response.Success(c, 201, "00000", "Workout created", data)
}

func (h *WorkoutHandler) UpdateWorkout(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.UpdateWorkout(id, req.Name, req.Date)
	return response.Success(c, 200, "00000", "Workout updated", data)
}

func (h *WorkoutHandler) DeleteWorkout(c echo.Context) error {
	data := h.provider.DeleteWorkout()
	return response.Success(c, 200, "00000", "Workout deleted", data)
}

func (h *WorkoutHandler) GetWorkoutStats(c echo.Context) error {
	data := h.provider.GetWorkoutStats()
	return response.Success(c, 200, "00000", "Success", data)
}
