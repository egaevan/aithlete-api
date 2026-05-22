package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	provider *mock.MockProvider
}

func NewProfileHandler(provider *mock.MockProvider) *ProfileHandler {
	return &ProfileHandler{provider: provider}
}

func (h *ProfileHandler) UpdateProfile(c echo.Context) error {
	var req request.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.UpdateProfile(req.Name, req.Birthday, req.Gender)
	return response.Success(c, 200, code.Success, message.MsgProfileUpdated, data)
}
