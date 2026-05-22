package service

import (
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type mockTokenService struct{}

func (m *mockTokenService) GenerateAccessToken(userID string) (string, int, error) {
	return "access_" + userID, 3600, nil
}

func (m *mockTokenService) GenerateRefreshToken(userID string) (string, error) {
	return "refresh_" + userID, nil
}

func (m *mockTokenService) ValidateAccessToken(token string) (string, error) {
	return "", nil
}

func (m *mockTokenService) ValidateRefreshToken(token string) (string, error) {
	return "", nil
}

type failingTokenService struct{}

func (m *failingTokenService) GenerateAccessToken(userID string) (string, int, error) {
	return "", 0, errors.New("token generation failed")
}

func (m *failingTokenService) GenerateRefreshToken(userID string) (string, error) {
	return "", errors.New("token generation failed")
}

func (m *failingTokenService) ValidateAccessToken(token string) (string, error) {
	return "", nil
}

func (m *failingTokenService) ValidateRefreshToken(token string) (string, error) {
	return "", nil
}

func TestGenerateAuthResult_Success(t *testing.T) {
	u := entity.NewUser("test@example.com", "Test User", "hashed:pass")
	u.SetID("user-123")

	result, err := GenerateAuthResult(u, &mockTokenService{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.User.ID != "user-123" {
		t.Errorf("expected ID user-123, got %s", result.User.ID)
	}
	if result.User.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", result.User.Email)
	}
	if result.Tokens.AccessToken != "access_user-123" {
		t.Errorf("expected access token access_user-123, got %s", result.Tokens.AccessToken)
	}
	if result.Tokens.RefreshToken != "refresh_user-123" {
		t.Errorf("expected refresh token refresh_user-123, got %s", result.Tokens.RefreshToken)
	}
	if result.Tokens.ExpiresIn != 3600 {
		t.Errorf("expected expiresIn 3600, got %d", result.Tokens.ExpiresIn)
	}
}

func TestGenerateAuthResult_AccessTokenFails(t *testing.T) {
	u := entity.NewUser("test@example.com", "Test User", "hashed:pass")
	u.SetID("user-123")

	_, err := GenerateAuthResult(u, &failingTokenService{})
	if err == nil {
		t.Fatal("expected error when access token generation fails")
	}
}

func TestIsNotFound_ReturnsTrueForErrUserNotFound(t *testing.T) {
	if !IsNotFound(domainerr.ErrUserNotFound) {
		t.Error("expected IsNotFound to return true for ErrUserNotFound")
	}
}

func TestIsNotFound_ReturnsFalseForOtherErrors(t *testing.T) {
	if IsNotFound(errors.New("other error")) {
		t.Error("expected IsNotFound to return false for other errors")
	}
}
