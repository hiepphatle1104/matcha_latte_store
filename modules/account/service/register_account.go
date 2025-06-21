package service

import (
	"context"
	"errors"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/dto"
)

func (s *AccountService) RegisterAccount(ctx context.Context, data *dto.SignUpAccount) error {
	exist, _ := s.repo.Find(ctx, map[string]string{"email": data.Email})
	if exist != nil {
		return errors.New("account already exists")
	}

	account, err := domain.NewAccount(data.Username, data.Email, data.Password)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, account)
}
