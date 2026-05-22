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

func (h *Handler) Register(c echo.Context) error {
	var req request.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.registerUseCase.Register(c.Request().Context(), req.Email, req.Name, req.Password)
	if err != nil {
		if errors.Is(err, domainerr.ErrEmailAlreadyExists) {
			return response.Error(c, http.StatusConflict, code.Conflict, message.MsgRegisterEmailExists)
		}
		return response.InternalServerError(c, message.MsgRegisterFailed)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgRegisterSuccess, toLoginResponse(result))
}
