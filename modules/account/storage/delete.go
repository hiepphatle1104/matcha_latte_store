package storage

import (
	"context"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain"
)

func (s *AccountStorage) Delete(ctx context.Context, filters map[string]string, data *domain.Account) error {
	return s.db.WithContext(ctx).Where(filters).Delete(data).Error
}
