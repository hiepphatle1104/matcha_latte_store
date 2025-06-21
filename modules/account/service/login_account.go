package service

import (
	"context"
	"errors"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/dto"
)

func (s *AccountService) LoginAccount(ctx context.Context, data *dto.SignInAccount) (string, error) {
	exist, _ := s.repo.Find(ctx, map[string]string{"email": data.Email})
	if exist == nil {
		return "", errors.New("account does not exist")
	}

	if err := exist.Password.Compare(data.Password); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := exist.AccountID.GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil
}
