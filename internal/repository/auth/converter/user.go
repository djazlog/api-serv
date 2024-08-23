package converter

import (
	"week/internal/model"
	modelRepo "week/internal/repository/auth/model"
)

func ToUserFromRepo(user *modelRepo.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
