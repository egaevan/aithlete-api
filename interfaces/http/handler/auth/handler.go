package auth

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	loginUseCase    usecase.LoginUseCase
	registerUseCase usecase.RegisterUseCase
	refreshUseCase  usecase.RefreshTokenUseCase
	getMeUseCase    usecase.GetMeUseCase
}

func New(login usecase.LoginUseCase, register usecase.RegisterUseCase, refresh usecase.RefreshTokenUseCase, getMe usecase.GetMeUseCase) *Handler {
	return &Handler{
		loginUseCase:    login,
		registerUseCase: register,
		refreshUseCase:  refresh,
		getMeUseCase:    getMe,
	}
}
