package dto

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email    string
	Name     string
	Password string
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type LoginResult struct {
	User   UserResult
	Tokens TokenResult
}

type UserResult struct {
	ID        string
	Email     string
	Name      string
	Avatar    string
	Birthday  string
	Gender    string
	CreatedAt string
	UpdatedAt string
}

type TokenResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

type AuthResponse struct {
	User   UserDTO
	Tokens TokenDTO
}

type UserDTO struct {
	ID        string
	Email     string
	Name      string
	Avatar    string
	Birthday  string
	Gender    string
	CreatedAt string
	UpdatedAt string
}

type TokenDTO struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

type ProfileUpdateRequest struct {
	Name     string
	Birthday string
	Gender   string
}
