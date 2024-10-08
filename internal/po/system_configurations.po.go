package po

import "gorm.io/gorm"

type SystemConfiguration struct {
	gorm.Model
	ConfigKey   string `json:"config_key"`
	ConfigValue string `json:"config_value"`
	Description string `json:"description"`
}

func (SystemConfiguration) TableName() string {
	return "system_configurations"
}
