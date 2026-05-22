package auth

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetMe(c echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return response.Error(c, http.StatusUnauthorized, code.Unauthorized, message.MsgUnauthorized)
	}

	result, err := h.getMeUseCase.GetMe(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, domainerr.ErrUserNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		return response.InternalServerError(c, message.MsgGetUserFailed)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toUserResponse(result))
}
