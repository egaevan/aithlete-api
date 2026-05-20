package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type ExerciseHandler struct {
	provider *mock.MockProvider
}

func NewExerciseHandler(provider *mock.MockProvider) *ExerciseHandler {
	return &ExerciseHandler{provider: provider}
}

func (h *ExerciseHandler) GetExercises(c echo.Context) error {
	data := h.provider.GetExercises()
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *ExerciseHandler) GetExercise(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.GetExercise(id)
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *ExerciseHandler) GetMuscleGroups(c echo.Context) error {
	data := h.provider.GetMuscleGroups()
	meta := response.NewMeta(len(data), 1, 20)
	return response.SuccessWithMeta(c, 200, "00000", "Success", data, meta)
}
