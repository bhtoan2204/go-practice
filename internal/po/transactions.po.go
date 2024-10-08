package po

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	AccountId       uint    `json:"account_id"`
	TransactionType string  `json:"transaction_type"` // 'deposit', 'withdrawal', 'transfer'
	Amount          float64 `json:"amount" gorm:"type:decimal(18,2)"`
	Currency        string  `json:"currency"`
	ReferenceId     string  `json:"reference_id"`
	Description     string  `json:"description"`
	Status          string  `json:"status"` // 'success', 'failed' default is 'pending'
}

func (Transaction) TableName() string {
	return "transactions"
}
