package workout

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

type updateSetRequestBody struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	RPE    int     `json:"rpe"`
}

func (h *Handler) UpdateSet(c echo.Context) error {
	var req updateSetRequestBody
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)
	workoutID := c.Param("id")
	exerciseID := c.Param("exerciseId")
	setID := c.Param("setId")

	result, err := h.updateSetUseCase.UpdateSet(c.Request().Context(), userID, workoutID, exerciseID, setID, req.Reps, req.Weight, req.RPE)
	if err != nil {
		if errors.Is(err, domainerr.ErrWorkoutNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		if errors.Is(err, domainerr.ErrExerciseNotFoundInWorkout) || errors.Is(err, domainerr.ErrSetNotFound) {
			return response.NotFound(c, err.Error())
		}
		if errors.Is(err, domainerr.ErrInvalidSetValue) {
			return response.BadRequest(c, err.Error())
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toWorkoutResponse(result))
}
