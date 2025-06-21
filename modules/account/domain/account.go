package domain

import (
	"github.com/google/uuid"
	obj "github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/value_obj"
)

type Account struct {
	AccountID obj.AccountID `json:"account_id" gorm:"column:account_id;"`
	Username  obj.Username  `json:"username" gorm:"column:username;"`
	Email     obj.Email     `json:"email" gorm:"column:email;"`
	Password  obj.Password  `json:"-" gorm:"column:password;"`
}

func (Account) TableName() string { return "accounts" }

func NewAccount(username, email, password string) (*Account, error) {
	account := &Account{
		AccountID: obj.AccountID(uuid.New().String()),
		Username:  obj.Username(username),
		Email:     obj.Email(email),
		Password:  obj.Password(password),
	}

	if err := account.AccountID.Validate(); err != nil {
		return nil, err
	}

	if err := account.Username.Validate(); err != nil {
		return nil, err
	}

	if err := account.Email.Validate(); err != nil {
		return nil, err
	}

	if err := account.Password.Hash(); err != nil {
		return nil, err
	}

	return account, nil
}
