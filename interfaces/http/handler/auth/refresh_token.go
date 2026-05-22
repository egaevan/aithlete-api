package auth

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RefreshToken(c echo.Context) error {
	var req request.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.refreshUseCase.RefreshToken(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgRefreshTokenExpired)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toTokenResponse(result))
}
