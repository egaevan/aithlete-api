package workout

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

func (h *Handler) Update(c echo.Context) error {
	var req request.UpdateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)
	id := c.Param("id")

	result, err := h.updateWorkoutUseCase.Update(c.Request().Context(), userID, id, req.Name, req.Date)
	if err != nil {
		if errors.Is(err, domainerr.ErrWorkoutNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusOK, code.Success, message.MsgWorkoutUpdated, toWorkoutResponse(result))
}
