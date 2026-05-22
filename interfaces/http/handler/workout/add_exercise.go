package workout

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

type addExerciseRequestBody struct {
	ExerciseID   string             `json:"exerciseId"`
	ExerciseName string             `json:"exerciseName"`
	Sets         []addSetRequestBody `json:"sets"`
}

type addSetRequestBody struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	RPE    int     `json:"rpe"`
}

func (h *Handler) AddExercise(c echo.Context) error {
	var req addExerciseRequestBody
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)
	id := c.Param("id")

	sets := make([]entity.Set, len(req.Sets))
	for i, s := range req.Sets {
		sets[i] = entity.Set{
			Reps:   s.Reps,
			Weight: s.Weight,
			RPE:    s.RPE,
		}
	}

	result, err := h.addExerciseUseCase.AddExercise(c.Request().Context(), userID, id,
		entity.ExerciseRef{ID: req.ExerciseID, Name: req.ExerciseName}, sets)
	if err != nil {
		if errors.Is(err, domainerr.ErrWorkoutNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		if errors.Is(err, domainerr.ErrDuplicateExercise) {
			return response.BadRequest(c, err.Error())
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toWorkoutResponse(result))
}
