package profile

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	updateProfileUseCase usecase.UpdateProfileUseCase
}

func New(updateProfile usecase.UpdateProfileUseCase) *Handler {
	return &Handler{
		updateProfileUseCase: updateProfile,
	}
}
