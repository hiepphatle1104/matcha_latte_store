package service

import (
	"context"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/dto"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/repository"
)

type AccountService struct {
	repo repository.AccountRepository
}

type IAccountService interface {
	RegisterAccount(context.Context, *dto.SignUpAccount) error
	LoginAccount(context.Context, *dto.SignInAccount) (string, error)
	GetAccount(context.Context, string) (*domain.Account, error)
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo}
}
