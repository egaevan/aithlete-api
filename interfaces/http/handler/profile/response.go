package profile

import "github.com/aithlete/aithlete-api/application/dto"

type ProfileResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func toProfileResponse(r *dto.UserResult) ProfileResponse {
	return ProfileResponse{
		ID:        r.ID,
		Email:     r.Email,
		Name:      r.Name,
		Avatar:    r.Avatar,
		Birthday:  r.Birthday,
		Gender:    r.Gender,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
