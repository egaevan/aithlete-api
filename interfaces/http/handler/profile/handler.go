package profile

import "github.com/aithlete/aithlete-api/application/usecase/profile"

type Handler struct {
	updateProfileUseCase profile.UpdateProfileUseCase
}

func New(updateProfile profile.UpdateProfileUseCase) *Handler {
	return &Handler{
		updateProfileUseCase: updateProfile,
	}
}
