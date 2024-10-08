package models

import (
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
