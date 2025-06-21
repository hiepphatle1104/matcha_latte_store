package value_obj

import (
	"database/sql/driver"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hiepphatle1104/matcha_latte_store/config"
	"time"
)

type AccountID string

func (a *AccountID) Validate() error {
	if len(*a) == 0 {
		return errors.New("AccountID cannot be empty")
	}
	_, err := uuid.Parse(string(*a))
	if err != nil {
		return errors.New("AccountID must be a valid UUID")
	}
	return nil
}

func (a *AccountID) Value() (driver.Value, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	return string(*a), nil
}

func (a *AccountID) Scan(value interface{}) error {
	if value == nil {
		*a = ""
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("AccountID must be a string")
	}

	*a = AccountID(str)
	return (*a).Validate()
}

func (a *AccountID) GenerateToken() (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "matcha_latte_store",
		Subject:   string(*a),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
