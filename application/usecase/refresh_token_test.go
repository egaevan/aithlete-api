package usecase

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/service"
)

func TestRefreshToken_Success(t *testing.T) {
	uc := NewRefreshTokenUseCase(&service.MockTokenService{})
	ctx := context.Background()

	result, err := uc.RefreshToken(ctx, "valid_refresh_token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.AccessToken == "" {
		t.Error("expected access token to be set")
	}
	if result.RefreshToken == "" {
		t.Error("expected refresh token to be set")
	}
	if result.ExpiresIn != 3600 {
		t.Errorf("expected expiresIn 3600, got %d", result.ExpiresIn)
	}
}

func TestRefreshToken_Invalid(t *testing.T) {
	uc := NewRefreshTokenUseCase(&service.MockTokenService{})
	ctx := context.Background()

	_, err := uc.RefreshToken(ctx, "invalid_refresh")
	if err == nil {
		t.Fatal("expected error for invalid refresh token")
	}
}
