package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Username  string    `gorm:"column:username;"`
	Password  string    `gorm:"column:password;"`
	Email     string    `gorm:"column:email;"`
	FirstName string    `gorm:"column:first_name;"`
	LastName  string    `gorm:"column:last_name;"`
	Status    string    `gorm:"column:status;"` // 'active', 'inactive', 'suspended'
	IsActive  bool      `gorm:"column:is_active;type:boolean"`
	Roles     []Role    `gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}
