package service

import (
	"context"
	"simple_bank/global"
	"simple_bank/internal/models"
	"simple_bank/internal/repository"
)

type AccountService interface {
	CRUDService[models.Account]
	ListByUserID(ctx context.Context, userID uint) ([]models.Account, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService() AccountService {
	return &accountService{accountRepo: repository.NewAccountRepository(global.MDB)}
}

func (s *accountService) Create(ctx context.Context, account *models.Account) error {
	return s.accountRepo.Create(ctx, account)
}

func (s *accountService) GetByID(ctx context.Context, id uint) (*models.Account, error) {
	return s.accountRepo.GetByID(ctx, id)
}

func (s *accountService) Update(ctx context.Context, account *models.Account) error {
	return s.accountRepo.Update(ctx, account)
}

func (s *accountService) Delete(ctx context.Context, id uint) error {
	return s.accountRepo.Delete(ctx, id)
}

func (s *accountService) ListByUserID(ctx context.Context, userID uint) ([]models.Account, error) {
	return s.accountRepo.ListByUserID(ctx, userID)
}
