package auth

import (
	"context"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/domain/service"
)

func TestGetMe_Success(t *testing.T) {
	userRepo := repository.NewMockUserRepository()
	getMeUC := NewGetMeUseCase(userRepo)
	regUC := NewRegisterUseCase(userRepo, &service.MockPasswordHasher{}, &service.MockTokenService{})

	ctx := context.Background()
	regResult, _ := regUC.Register(ctx, "test@example.com", "Test User", "password123")

	result, err := getMeUC.GetMe(ctx, regResult.User.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", result.Email)
	}
	if result.Name != "Test User" {
		t.Errorf("expected name Test User, got %s", result.Name)
	}
}

func TestGetMe_UserNotFound(t *testing.T) {
	getMeUC := NewGetMeUseCase(repository.NewMockUserRepository())
	ctx := context.Background()

	_, err := getMeUC.GetMe(ctx, "nonexistent-id")
	if err == nil {
		t.Fatal("expected error for nonexistent user")
	}
}
