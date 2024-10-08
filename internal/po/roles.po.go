package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          int64        `gorm:"column:id;primaryKey;type:int;not null;unique;index:autoIncrement;"`
	Name        string       `gorm:"column:name;"`
	Description string       `gorm:"column:description;type:text;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (Role) TableName() string {
	return "roles"
}
