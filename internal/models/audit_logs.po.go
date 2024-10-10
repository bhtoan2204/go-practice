package models

type AuditLog struct {
	BaseModel
	UserID      string `json:"user_id"`
	Action      string `json:"action"`
	Details     string `json:"details"`
	IpAddress   string `json:"ip_address"`
	DeviceToken string `json:"device_token"`
}
