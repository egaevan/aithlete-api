package auth

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

func (h *Handler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.loginUseCase.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domainerr.ErrInvalidCredentials) {
			return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgInvalidCredentials)
		}
		return response.InternalServerError(c, message.MsgLoginFailed)
	}

	return response.Success(c, http.StatusOK, code.Success, message.MsgLoginSuccess, toLoginResponse(result))
}
