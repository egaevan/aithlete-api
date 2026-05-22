package service

type TokenService interface {
	GenerateAccessToken(userID string) (string, int, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, error)
}
