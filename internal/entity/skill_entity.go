package entity

import "github.com/google/uuid"

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SkillName string    `gorm:"column:skill_name;"`
}
