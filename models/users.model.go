package models

import (
	"time"
)

type User struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username          string    `gorm:"size:50;not null;unique" json:"username"`
	Email             string    `gorm:"size:100;not null;unique" json:"email"`
	Password          string    `gorm:"size:255;not null" json:"password"`
	Firstname         string    `gorm:"size:100" json:"firstname"`
	Lastname          string    `gorm:"size:100" json:"lastname"`
	Bio               string    `gorm:"type:text" json:"bio"`
	ProfilePictureURL string    `gorm:"size:255" json:"profile_picture_url"`
	IsAdmin           bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
