package po

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PermissionName string `json:"permission_name"`
	Description    string `json:"description"`
	Roles          []Role `gorm:"many2many:role_permissions;"`
}

func (Permission) TableName() string {
	return "permissions"
}
