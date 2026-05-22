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

func (h *Handler) Create(c echo.Context) error {
	var req request.CreateWorkoutRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)

	result, err := h.createWorkoutUseCase.Create(c.Request().Context(), userID, req.Name, req.Date, req.WeightUnit, req.Notes)
	if err != nil {
		if errors.Is(err, domainerr.ErrEmptyWorkout) {
			return response.BadRequest(c, err.Error())
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgWorkoutCreated, toWorkoutResponse(result))
}
