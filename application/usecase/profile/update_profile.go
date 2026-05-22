package profile

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/application/service"
	"github.com/aithlete/aithlete-api/domain/repository"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

type UpdateProfileUseCase interface {
	UpdateProfile(ctx context.Context, userID, name, birthday, gender string) (*dto.UserResult, error)
}

type updateProfileUseCase struct {
	userRepo repository.UserRepository
}

func NewUpdateProfileUseCase(userRepo repository.UserRepository) UpdateProfileUseCase {
	return &updateProfileUseCase{userRepo: userRepo}
}

func (uc *updateProfileUseCase) UpdateProfile(ctx context.Context, userID, name, birthday, gender string) (*dto.UserResult, error) {
	u, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		if service.IsNotFound(err) {
			return nil, domainerr.ErrUserNotFound
		}
		return nil, fmt.Errorf("find user: %w", err)
	}

	u.UpdateProfile(name, birthday, gender)

	if err := uc.userRepo.Update(ctx, u); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return mapper.UserToResult(u), nil
}
