package repository

import (
	"context"
	"simple_bank/internal/models"

	"gorm.io/gorm"
)

type AccountRepository interface {
	CRUDRepository[models.Account]
	ListByUserID(ctx context.Context, userID uint) ([]models.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

// ListByUserID implements AccountRepository.
func (r *accountRepository) ListByUserID(ctx context.Context, userID uint) ([]models.Account, error) {
	panic("unimplemented")
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(ctx context.Context, account *models.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) GetByID(ctx context.Context, id uint) (*models.Account, error) {
	var account models.Account
	err := r.db.WithContext(ctx).First(&account, id).Error
	return &account, err
}

func (r *accountRepository) Update(ctx context.Context, account *models.Account) error {
	return r.db.WithContext(ctx).Save(account).Error
}

func (r *accountRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Account{}, id).Error
}
