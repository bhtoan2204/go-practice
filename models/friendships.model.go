package models

import (
	"time"
)

type FriendshipStatus string

const (
	Pending  FriendshipStatus = "pending"
	Accepted FriendshipStatus = "accepted"
	Declined FriendshipStatus = "declined"
	Blocked  FriendshipStatus = "blocked"
)

type Friendship struct {
	ID        uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint             `gorm:"not null" json:"user_id"`
	FriendID  uint             `gorm:"not null" json:"friend_id"`
	Status    FriendshipStatus `gorm:"type:enum('pending', 'accepted', 'declined', 'blocked');not null;default:'pending'" json:"status"`
	CreatedAt time.Time        `gorm:"autoCreateTime" json:"created_at"`
}
