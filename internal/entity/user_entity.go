package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"column:name;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	Bio       string    `gorm:"column:bio;"`
	Address   string    `gorm:"column:address;"`
	Role      string    `gorm:"column:role;"`
	AvatarUrl string    `gorm:"column:avatar_url;"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}
