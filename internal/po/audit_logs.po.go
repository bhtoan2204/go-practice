package po

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"column:user_id"`
	Action      string `json:"action"`
	Details     string `json:"details"`
	IpAddress   string `json:"ip_address" gorm:"size:45"` // For IPv4/IPv6
	DeviceToken string `json:"device_token"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
