package storage

import (
	"context"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
)

func (s *AccountStorage) Save(ctx context.Context, data *domain.Account) error {
	return s.db.WithContext(ctx).Create(data).Error
}
