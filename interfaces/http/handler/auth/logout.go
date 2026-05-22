package auth

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Logout(c echo.Context) error {
	return response.Success(c, http.StatusOK, code.Success, message.MsgLogoutSuccess, map[string]any{})
}
