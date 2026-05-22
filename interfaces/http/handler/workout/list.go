package workout

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) List(c echo.Context) error {
	userID := c.Get("user_id").(string)

	results, err := h.listWorkoutsUseCase.List(c.Request().Context(), userID)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	resp := make([]WorkoutResponse, len(results))
	for i, r := range results {
		resp[i] = toWorkoutResponse(&r)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, resp)
}
