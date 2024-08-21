package models

import (
	"time"
)

type NotificationType string

const (
	LikeNotification          NotificationType = "like"
	CommentNotification       NotificationType = "comment"
	FriendRequestNotification NotificationType = "friend_request"
	MessageNotification       NotificationType = "message"
)

type Notification struct {
	ID          uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint             `gorm:"not null" json:"user_id"`
	Type        NotificationType `gorm:"type:enum('like', 'comment', 'friend_request', 'message');not null" json:"type"`
	ReferenceID uint             `gorm:"not null" json:"reference_id"`
	Content     string           `gorm:"type:text" json:"content"`
	IsRead      bool             `gorm:"not null;default:false" json:"is_read"`
	CreatedAt   time.Time        `gorm:"autoCreateTime" json:"created_at"`
}
