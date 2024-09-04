package models

import (
	"time"
)

type GroupRole string

const (
	MemberRole GroupRole = "member"
	AdminRole  GroupRole = "admin"
)

type GroupMember struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID  uint      `gorm:"not null" json:"group_id"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Role     GroupRole `gorm:"type:enum('member', 'admin');not null;default:'member'" json:"role"`
	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`
}
