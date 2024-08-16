package models

import (
	"time"
)

type Message struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SenderID   uint      `gorm:"not null" json:"sender_id"`
	ReceiverID uint      `gorm:"not null" json:"receiver_id"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	IsRead     bool      `gorm:"not null;default:false" json:"is_read"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
