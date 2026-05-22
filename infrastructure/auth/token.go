package auth

import (
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/service"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret     []byte
	expiration time.Duration
}

func NewTokenService(secret string, expirationHours int) *TokenService {
	return &TokenService{
		secret:     []byte(secret),
		expiration: time.Duration(expirationHours) * time.Hour,
	}
}

var _ service.TokenService = (*TokenService)(nil)

type claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *TokenService) GenerateAccessToken(userID string) (string, int, error) {
	now := time.Now()
	exp := now.Add(s.expiration)

	c := claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "aithlete-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", 0, fmt.Errorf("sign token: %w", err)
	}

	expiresIn := int(time.Until(exp).Seconds())
	return signed, expiresIn, nil
}

func (s *TokenService) GenerateRefreshToken(userID string) (string, error) {
	now := time.Now()
	exp := now.Add(30 * 24 * time.Hour)

	c := claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "aithlete-api",
			ID:        "refresh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("sign refresh token: %w", err)
	}

	return signed, nil
}

func (s *TokenService) ValidateAccessToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return "", fmt.Errorf("parse token: %w", err)
	}

	c, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return c.UserID, nil
}

func (s *TokenService) ValidateRefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return "", fmt.Errorf("parse refresh token: %w", err)
	}

	c, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid refresh token")
	}

	if c.ID != "refresh" {
		return "", fmt.Errorf("not a refresh token")
	}

	return c.UserID, nil
}
