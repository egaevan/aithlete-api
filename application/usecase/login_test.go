package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestLogin_Success(t *testing.T) {
	userRepo := repository.NewMockUserRepository()
	loginUC := NewLoginUseCase(userRepo, &service.MockPasswordHasher{}, &service.MockTokenService{})

	u := entity.NewUser("test@example.com", "Test User", "hashed:password123")
	u.SetID("user-test")
	userRepo.Users[u.GetEmail()] = u

	ctx := context.Background()
	result, err := loginUC.Login(ctx, "test@example.com", "password123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.User.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", result.User.Email)
	}
	if result.Tokens.AccessToken == "" {
		t.Error("expected access token to be set")
	}
}

func TestLogin_InvalidEmail(t *testing.T) {
	loginUC := NewLoginUseCase(repository.NewMockUserRepository(), &service.MockPasswordHasher{}, &service.MockTokenService{})
	ctx := context.Background()

	_, err := loginUC.Login(ctx, "nonexistent@example.com", "password123")
	if err == nil {
		t.Fatal("expected error for invalid email")
	}
		if !errors.Is(err, domainerr.ErrInvalidCredentials) {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	userRepo := repository.NewMockUserRepository()
	loginUC := NewLoginUseCase(userRepo, &service.MockPasswordHasher{}, &service.MockTokenService{})

	u := entity.NewUser("test@example.com", "Test User", "hashed:password123")
	u.SetID("user-test")
	userRepo.Users[u.GetEmail()] = u

	ctx := context.Background()
	_, err := loginUC.Login(ctx, "test@example.com", "wrongpassword")
	if err == nil {
		t.Fatal("expected error for wrong password")
	}
		if !errors.Is(err, domainerr.ErrInvalidCredentials) {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}
