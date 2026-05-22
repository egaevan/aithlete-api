package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestRegister_Success(t *testing.T) {
	uc := NewRegisterUseCase(repository.NewMockUserRepository(), &service.MockPasswordHasher{}, &service.MockTokenService{})
	ctx := context.Background()

	result, err := uc.Register(ctx, "test@example.com", "Test User", "password123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.User.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", result.User.Email)
	}
	if result.User.Name != "Test User" {
		t.Errorf("expected name Test User, got %s", result.User.Name)
	}
	if result.User.ID == "" {
		t.Error("expected user ID to be set")
	}
	if result.Tokens.AccessToken == "" {
		t.Error("expected access token to be set")
	}
	if result.Tokens.RefreshToken == "" {
		t.Error("expected refresh token to be set")
	}
	if result.Tokens.ExpiresIn != 3600 {
		t.Errorf("expected expiresIn 3600, got %d", result.Tokens.ExpiresIn)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	uc := NewRegisterUseCase(repository.NewMockUserRepository(), &service.MockPasswordHasher{}, &service.MockTokenService{})
	ctx := context.Background()

	uc.Register(ctx, "test@example.com", "Test User", "password123")

	_, err := uc.Register(ctx, "test@example.com", "Another User", "password456")
	if err == nil {
		t.Fatal("expected error for duplicate email")
	}
		if !errors.Is(err, domainerr.ErrEmailAlreadyExists) {
		t.Errorf("expected ErrEmailAlreadyExists, got %v", err)
	}
}
