package models

import (
	"errors"
	"strings"
)

type Account struct {
	BaseModel
	UserID        string `json:"user_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	Balance       int64  `json:"balance"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
}

func isAlphaNum(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func (a *Account) Validate() error {
	var validationErrors []string

	if a.UserID == "" {
		validationErrors = append(validationErrors, "User ID is required.")
	}

	if strings.TrimSpace(a.AccountNumber) == "" {
		validationErrors = append(validationErrors, "Account number is required.")
	} else if !isAlphaNum(a.AccountNumber) {
		validationErrors = append(validationErrors, "Account number must be alphanumeric.")
	}

	validAccountTypes := map[string]bool{
		"savings":  true,
		"checking": true,
	}
	if !validAccountTypes[a.AccountType] {
		validationErrors = append(validationErrors, "Account type must be either 'savings' or 'checking'.")
	}

	if a.Balance < 0 {
		validationErrors = append(validationErrors, "Balance must be greater than or equal to 0.")
	}

	if len(a.Currency) != 3 {
		validationErrors = append(validationErrors, "Currency must be exactly 3 characters.")
	}

	validStatuses := map[string]bool{
		"active":    true,
		"closed":    true,
		"suspended": true,
	}
	if !validStatuses[a.Status] {
		validationErrors = append(validationErrors, "Status must be one of 'active', 'closed', or 'suspended'.")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

func NewAccount(userID string, accountNumber, accountType, currency string, balance int64) (*Account, error) {
	account := &Account{
		UserID:        userID,
		AccountNumber: accountNumber,
		AccountType:   accountType,
		Currency:      currency,
		Balance:       balance,
		Status:        "active",
	}

	// Thực hiện validation
	if err := account.Validate(); err != nil {
		return nil, err
	}

	return account, nil
}
