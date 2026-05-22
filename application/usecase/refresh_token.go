package usecase

import (
	"context"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/service"
)

type RefreshTokenUseCase interface {
	RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenResult, error)
}

type refreshTokenUseCase struct {
	tokens service.TokenService
}

func NewRefreshTokenUseCase(tokens service.TokenService) RefreshTokenUseCase {
	return &refreshTokenUseCase{
		tokens: tokens,
	}
}

func (uc *refreshTokenUseCase) RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenResult, error) {
	userID, err := uc.tokens.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	accessToken, expiresIn, err := uc.tokens.GenerateAccessToken(userID)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	newRefreshToken, err := uc.tokens.GenerateRefreshToken(userID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &dto.TokenResult{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}
