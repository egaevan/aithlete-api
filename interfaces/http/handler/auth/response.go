package auth

import "github.com/aithlete/aithlete-api/application/dto"

type LoginResponse struct {
	User   UserResponse  `json:"user"`
	Tokens TokenResponse `json:"tokens"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Birthday  string `json:"birthday,omitempty"`
	Gender    string `json:"gender,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

func toLoginResponse(r *dto.LoginResult) LoginResponse {
	return LoginResponse{
		User:   toUserResponse(&r.User),
		Tokens: toTokenResponse(&r.Tokens),
	}
}

func toUserResponse(u *dto.UserResult) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Avatar:    u.Avatar,
		Birthday:  u.Birthday,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func toTokenResponse(t *dto.TokenResult) TokenResponse {
	return TokenResponse{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		ExpiresIn:    t.ExpiresIn,
	}
}
