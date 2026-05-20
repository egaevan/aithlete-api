package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	provider *mock.MockProvider
}

func NewAuthHandler(provider *mock.MockProvider) *AuthHandler {
	return &AuthHandler{provider: provider}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.Login(req.Email, req.Password)
	return response.Success(c, 200, "00000", "Login successful", data)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req request.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.Register(req.Email, req.Name, req.Password)
	return response.Success(c, 201, "00000", "Account created successfully", data)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	data := h.provider.Logout()
	return response.Success(c, 200, "00000", "Logged out successfully", data)
}

func (h *AuthHandler) GetMe(c echo.Context) error {
	data := h.provider.GetMe()
	return response.Success(c, 200, "00000", "Success", data)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var req request.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	data := h.provider.RefreshToken()
	return response.Success(c, 200, "00000", "Success", data)
}
