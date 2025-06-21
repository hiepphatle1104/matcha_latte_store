package service

import (
	"context"
	"errors"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
	"strings"
)

func (s *AccountService) GetAccount(ctx context.Context, accountID string) (*domain.Account, error) {
	if strings.TrimSpace(accountID) == "" {
		return nil, errors.New("accountID is empty")
	}

	return s.repo.Find(ctx, map[string]string{"account_id": accountID})
}
