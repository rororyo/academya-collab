package model

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Bio       string     `json:"bio,omitempty"`
	Address   string     `json:"address,omitempty"`
	Role      string     `json:"role,omitempty" validate:"required,oneof=admin user"`
	AvatarUrl string     `json:"avatar_url,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=255" json:"token"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type SearchUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Page  int    `json:"page,omitempty" validate:"min=1"`
	Size  int    `json:"size,omitempty" validate:"min=1,max=100"`
}

type UpdateUserRequest struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Address   string `json:"address,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type DeleteUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
