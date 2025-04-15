package entity

import "github.com/google/uuid"

type JobSkills struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	JobID   uuid.UUID `gorm:"column:job_id;not null"`
	SkillID uuid.UUID `gorm:"column:skill_id;not null"`
	Job     Job       `gorm:"foreignKey:JobID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Skill   Skill     `gorm:"foreignKey:SkillID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
