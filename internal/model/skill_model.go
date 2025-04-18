package model

import "github.com/google/uuid"

type SkillResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	SkillName string     `json:"skill_name,omitempty"`
}

type SkillRequest struct {
	SkillName string `json:"skill_name" validate:"required,max=100"`
}

type SearchSkillRequest struct {
	SkillName string `json:"skill_name,omitempty"`
	Page      int    `json:"page,omitempty" validate:"min=1"`
	Size      int    `json:"size,omitempty" validate:"min=1,max=100"`
}

type GetSkillRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type UpdateSkillRequest struct {
	ID        string `json:"-" validate:"required,max=100"`
	SkillName string `json:"skill_name,omitempty"`
}

type DeleteSkillRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
