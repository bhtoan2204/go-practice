package po

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
