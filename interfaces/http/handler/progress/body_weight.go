package progress

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBodyWeightHistory(c echo.Context) error {
	userID := c.Get("user_id").(string)

	results, err := h.getBodyWeightHistoryUseCase.GetBodyWeightHistory(c.Request().Context(), userID)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toBodyWeightResponseList(results))
}

func (h *Handler) AddBodyWeight(c echo.Context) error {
	var req request.AddBodyWeightRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)

	result, err := h.addBodyWeightUseCase.Add(c.Request().Context(), userID, "", req.Weight, req.BodyFatPercentage)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgBodyWeightAdded, toBodyWeightResponse(result))
}
