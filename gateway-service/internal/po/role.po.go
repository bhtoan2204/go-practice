package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          int64  `gorm:"column:id;primaryKey;type:int;not null;unique;index:autoIncrement;"`
	Name        string `gorm:"column:name;"`
	Description string `gorm:"column:description;type:text;"`
}

func (Role) TableName() string {
	return "hr_roles"
}
