package exercise

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Get(c echo.Context) error {
	id := c.Param("id")

	result, err := h.getExerciseUseCase.Get(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, domainerr.ErrExerciseNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toExerciseResponse(result))
}
