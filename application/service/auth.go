package service

import (
	"errors"
	"fmt"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/application/mapper"
	"github.com/aithlete/aithlete-api/domain/entity"
	domainservice "github.com/aithlete/aithlete-api/domain/service"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
)

func GenerateAuthResult(u *entity.User, tokens domainservice.TokenService) (*dto.LoginResult, error) {
	accessToken, expiresIn, err := tokens.GenerateAccessToken(u.GetID())
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := tokens.GenerateRefreshToken(u.GetID())
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &dto.LoginResult{
		User: *mapper.UserToResult(u),
		Tokens: dto.TokenResult{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    expiresIn,
		},
	}, nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, domainerr.ErrUserNotFound)
}
