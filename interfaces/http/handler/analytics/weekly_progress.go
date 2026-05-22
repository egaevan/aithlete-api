package analytics

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetWeeklyProgress(c echo.Context) error {
	userID := c.Get("user_id").(string)

	results, err := h.weeklyProgressUseCase.Get(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, domainerr.ErrNoAnalyticsData) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toWeeklyProgressDayResponseList(results))
}
