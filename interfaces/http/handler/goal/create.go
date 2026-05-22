package goal

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var req request.CreateGoalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.createGoalUseCase.Create(c.Request().Context(), userID, req.Title, req.Type, req.Target, req.Unit, req.Period, req.Deadline)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgCreated, toGoalResponse(result))
}
