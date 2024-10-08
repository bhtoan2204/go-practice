package po

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"column:user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	IsRead  bool   `json:"is_read"`
}

func (Notification) TableName() string {
	return "notifications"
}
