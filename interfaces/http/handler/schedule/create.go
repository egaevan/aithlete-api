package schedule

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req request.CreateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)

	result, err := h.createScheduleUseCase.Create(c.Request().Context(), userID, req.Date, req.Time, req.Title, req.Duration, req.Type, req.Notes)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgCreated, toScheduleResponse(result))
}
