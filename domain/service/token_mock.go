package service

type MockTokenService struct{}

func (m *MockTokenService) GenerateAccessToken(userID string) (string, int, error) {
	return "access_" + userID, 3600, nil
}

func (m *MockTokenService) GenerateRefreshToken(userID string) (string, error) {
	return "refresh_" + userID, nil
}

func (m *MockTokenService) ValidateAccessToken(token string) (string, error) {
	return "", nil
}

func (m *MockTokenService) ValidateRefreshToken(token string) (string, error) {
	if token == "invalid_refresh" {
		return "", &mockTokenError{"invalid token"}
	}
	return "user-" + token, nil
}

type mockTokenError struct{ msg string }

func (e *mockTokenError) Error() string { return e.msg }
