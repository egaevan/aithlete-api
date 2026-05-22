package auth

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	domainservice "github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type LoginUseCase interface {
	Login(ctx context.Context, email, password string) (*dto.LoginResult, error)
}

type loginUseCase struct {
	userRepo repository.UserRepository
	hasher   domainservice.PasswordHasher
	tokens   domainservice.TokenService
}

func NewLoginUseCase(userRepo repository.UserRepository, hasher domainservice.PasswordHasher, tokens domainservice.TokenService) LoginUseCase {
	return &loginUseCase{
		userRepo: userRepo,
		hasher:   hasher,
		tokens:   tokens,
	}
}

func (uc *loginUseCase) Login(ctx context.Context, email, password string) (*dto.LoginResult, error) {
	u, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if service.IsNotFound(err) {
			return nil, domainerr.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("find user: %w", err)
	}

	if !uc.hasher.Verify(password, u.PasswordHash()) {
		return nil, domainerr.ErrInvalidCredentials
	}

	return service.GenerateAuthResult(u, uc.tokens)
}
