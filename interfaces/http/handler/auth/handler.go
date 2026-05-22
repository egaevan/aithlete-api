package auth

import "github.com/aithlete/aithlete-api/application/usecase/auth"

type Handler struct {
	loginUseCase    auth.LoginUseCase
	registerUseCase auth.RegisterUseCase
	refreshUseCase  auth.RefreshTokenUseCase
	getMeUseCase    auth.GetMeUseCase
}

func New(login auth.LoginUseCase, register auth.RegisterUseCase, refresh auth.RefreshTokenUseCase, getMe auth.GetMeUseCase) *Handler {
	return &Handler{
		loginUseCase:    login,
		registerUseCase: register,
		refreshUseCase:  refresh,
		getMeUseCase:    getMe,
	}
}
