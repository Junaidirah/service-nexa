package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Sessions  []Session      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
	RefreshToken string    `gorm:"type:text;not null;index"`
	UserAgent    string    `gorm:"type:varchar(255)"`
	ClientIP     string    `gorm:"type:varchar(45)"`
	IsBlocked    bool      `gorm:"default:false"`
	ExpiresAt    time.Time `gorm:"not null;index"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func EnableUUIDExtension(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
}

func (User) TableName() string {
	return "users"
}

func (Session) TableName() string {
	return "sessions"
}
