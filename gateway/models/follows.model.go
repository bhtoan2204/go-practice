package models

import (
	"time"
)

type Follow struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FollowerID  uint      `gorm:"not null" json:"follower_id"`
	FollowingID uint      `gorm:"not null" json:"following_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
