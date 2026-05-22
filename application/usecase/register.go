package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/domain/repository"
	domainservice "github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type RegisterUseCase interface {
	Register(ctx context.Context, email, name, password string) (*dto.LoginResult, error)
}

type registerUseCase struct {
	userRepo repository.UserRepository
	hasher   domainservice.PasswordHasher
	tokens   domainservice.TokenService
}

func NewRegisterUseCase(userRepo repository.UserRepository, hasher domainservice.PasswordHasher, tokens domainservice.TokenService) RegisterUseCase {
	return &registerUseCase{
		userRepo: userRepo,
		hasher:   hasher,
		tokens:   tokens,
	}
}

func (uc *registerUseCase) Register(ctx context.Context, email, name, password string) (*dto.LoginResult, error) {
	existing, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil && !service.IsNotFound(err) {
		return nil, fmt.Errorf("check existing user: %w", err)
	}
	if existing != nil {
		return nil, domainerr.ErrEmailAlreadyExists
	}

	hashed, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	u := entity.NewUser(email, name, hashed)
	if err := uc.userRepo.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return service.GenerateAuthResult(u, uc.tokens)
}
