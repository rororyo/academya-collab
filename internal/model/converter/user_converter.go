package converter

import (
	"collab-be/internal/entity"
	"collab-be/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        &user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Address:   user.Address,
		AvatarUrl: user.AvatarUrl,
		Role:      user.Role,
		Token:     user.Token,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}
