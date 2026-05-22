package profile

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

func (h *Handler) UpdateProfile(c echo.Context) error {
	var req request.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	userID := c.Get("user_id").(string)

	result, err := h.updateProfileUseCase.UpdateProfile(c.Request().Context(), userID, req.Name, req.Birthday, req.Gender)
	if err != nil {
		if errors.Is(err, domainerr.ErrUserNotFound) {
			return response.NotFound(c, message.MsgUserNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusOK, code.Success, message.MsgProfileUpdated, toProfileResponse(result))
}
