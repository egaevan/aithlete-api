package goal

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateProgress(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	var req request.UpdateGoalProgressRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.updateGoalProgressUseCase.UpdateProgress(c.Request().Context(), userID, goalID, req.Current)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponse(result))
}
