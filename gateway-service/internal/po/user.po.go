package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;primaryKey;type:varchar(255);not null;unique;index:idx_uuid;"`
	Username string    `gorm:"column:username;"`
	Password string    `gorm:"column:password;"`
	Roles    []Role    `gorm:"many2many:hr_user_roles;"`
	IsActive bool      `gorm:"column:is_active;type:boolean"`
}

func (User) TableName() string {
	return "hr_users"
}
