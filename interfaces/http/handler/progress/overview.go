package progress

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProgressOverview(c echo.Context) error {
	userID := c.Get("user_id").(string)

	result, err := h.getProgressOverviewUseCase.GetOverview(c.Request().Context(), userID)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toProgressOverviewResponse(result))
}
