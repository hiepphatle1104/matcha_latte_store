package storage

import (
	"gorm.io/gorm"
)

type AccountStorage struct {
	db *gorm.DB
}

func NewAccountStorage(db *gorm.DB) *AccountStorage {
	return &AccountStorage{db}
}
