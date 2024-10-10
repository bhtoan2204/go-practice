package models

type Notification struct {
	BaseModel
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	IsRead  bool   `json:"is_read"`
}
