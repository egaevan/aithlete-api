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

func (h *Handler) Complete(c echo.Context) error {
	userID := c.Get("user_id").(string)
	id := c.Param("id")

	result, err := h.completeWorkoutUseCase.Complete(c.Request().Context(), userID, id)
	if err != nil {
		if errors.Is(err, domainerr.ErrWorkoutNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		if errors.Is(err, domainerr.ErrEmptyWorkout) || errors.Is(err, domainerr.ErrIncompleteWorkout) {
			return response.BadRequest(c, err.Error())
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toWorkoutResponse(result))
}
