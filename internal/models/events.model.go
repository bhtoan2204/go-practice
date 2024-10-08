package models

type Event struct {
	BaseModel
	EventType    string `json:"event_type"`
	EventPayload any    `json:"event_data" gorm:"type:json"`
	Status       string `json:"status"`
}
