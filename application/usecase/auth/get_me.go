package auth

import (
	"context"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/repository"
)

type GetMeUseCase interface {
	GetMe(ctx context.Context, userID string) (*dto.UserResult, error)
}

type getMeUseCase struct {
	userRepo repository.UserRepository
}

func NewGetMeUseCase(userRepo repository.UserRepository) GetMeUseCase {
	return &getMeUseCase{
		userRepo: userRepo,
	}
}

func (uc *getMeUseCase) GetMe(ctx context.Context, userID string) (*dto.UserResult, error) {
	u, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.UserToResult(u), nil
}
