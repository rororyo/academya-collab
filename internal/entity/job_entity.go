package entity

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CompanyID   uuid.UUID `gorm:"column:recruiter_id;not null"`
	Title       string    `gorm:"column:title;not null"`
	Position    string    `gorm:"column:position;not null"`
	Description string    `gorm:"column:description;not null"`
	Location    string    `gorm:"column:location;not null"`
	Salary      int       `gorm:"column:salary"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Foreign key relation
	Company User `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
