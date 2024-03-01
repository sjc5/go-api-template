package model

import (
	"time"
)

type User struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	Email           string    `gorm:"unique" json:"email"`
	EmailIsVerified bool      `gorm:"not null; default:false" json:"email_is_verified"`
	CreatedAt       time.Time `gorm:"not null; default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null; default:current_timestamp; autoUpdateTime" json:"updated_at"`
	Sessions        []Session `gorm:"foreignKey:UserID"`
}

type Session struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"not null; default:current_timestamp" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
}
