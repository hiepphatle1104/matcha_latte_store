package repository

import (
	"context"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
)

type AccountRepository interface {
	Save(context.Context, *domain.Account) error
	Find(context.Context, map[string]string) (*domain.Account, error)
	Delete(context.Context, map[string]string, *domain.Account) error
}
