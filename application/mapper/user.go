package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func UserToResult(u *entity.User) *dto.UserResult {
	return &dto.UserResult{
		ID:        u.GetID(),
		Email:     u.GetEmail(),
		Name:      u.GetName(),
		Avatar:    u.GetAvatar(),
		Birthday:  u.GetBirthday(),
		Gender:    u.GetGender(),
		CreatedAt: u.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: u.GetUpdatedAt().Format(time.RFC3339),
	}
}
