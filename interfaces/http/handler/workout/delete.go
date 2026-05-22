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

func (h *Handler) Delete(c echo.Context) error {
	userID := c.Get("user_id").(string)
	id := c.Param("id")

	err := h.deleteWorkoutUseCase.Delete(c.Request().Context(), userID, id)
	if err != nil {
		if errors.Is(err, domainerr.ErrWorkoutNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusOK, code.Success, message.MsgWorkoutDeleted, map[string]any{})
}
