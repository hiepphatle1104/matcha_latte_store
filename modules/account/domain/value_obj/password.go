package value_obj

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p *Password) Validate() error {
	if len(*p) == 0 {
		return errors.New("password cannot be empty")
	}

	if len(*p) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}

func (p *Password) Compare(plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(*p), []byte(plainPassword))
}

func (p *Password) Hash() error {
	if err := p.Validate(); err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*p = Password(hashedPass)
	return nil
}

func (p *Password) Value() (string, error) {
	if err := p.Validate(); err != nil {
		return "", err
	}

	return string(*p), nil
}

func (p *Password) Scan(value interface{}) error {
	if value == nil {
		*p = ""
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("password must be a string")
	}

	*p = Password(str)
	return nil
}
