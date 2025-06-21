package storage

import (
	"context"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
)

func (s *AccountStorage) Find(ctx context.Context, filters map[string]string) (*domain.Account, error) {
	var account *domain.Account
	if err := s.db.WithContext(ctx).Where(filters).First(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}
