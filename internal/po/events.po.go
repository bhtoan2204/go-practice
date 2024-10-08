package po

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventType    string `json:"event_type"`
	EventPayload any    `json:"event_data" gorm:"type:json"`
	Status       string `json:"status"`
}

func (Event) TableName() string {
	return "events"
}
