package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func TestUpdateProfile_Success(t *testing.T) {
	userRepo := repository.NewMockUserRepository()
	regUC := NewRegisterUseCase(userRepo, &service.MockPasswordHasher{}, &service.MockTokenService{})
	profileUC := NewUpdateProfileUseCase(userRepo)
	ctx := context.Background()

	regResult, _ := regUC.Register(ctx, "test@example.com", "Test User", "password123")

	result, err := profileUC.UpdateProfile(ctx, regResult.User.ID, "New Name", "1995-06-15", "male")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "New Name" {
		t.Errorf("expected name 'New Name', got %s", result.Name)
	}
	if result.Birthday != "1995-06-15" {
		t.Errorf("expected birthday '1995-06-15', got %s", result.Birthday)
	}
	if result.Gender != "male" {
		t.Errorf("expected gender 'male', got %s", result.Gender)
	}
	if result.Email != "test@example.com" {
		t.Errorf("expected email to remain unchanged, got %s", result.Email)
	}
}

func TestUpdateProfile_UserNotFound(t *testing.T) {
	profileUC := NewUpdateProfileUseCase(repository.NewMockUserRepository())
	ctx := context.Background()

	_, err := profileUC.UpdateProfile(ctx, "nonexistent-id", "New Name", "1995-06-15", "male")
	if err == nil {
		t.Fatal("expected error for nonexistent user")
	}
	if !errors.Is(err, domainerr.ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}
