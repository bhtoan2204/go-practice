package po

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID        uint   `json:"user_id" gorm:"column:user_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"` // 'savings', 'checking'
	Balance       int64  `json:"balance"`
	Currency      string `json:"currency"`
	Status        string `json:"status"` // 'active', 'closed', 'suspended'
}

func (Account) TableName() string {
	return "accounts"
}
