package value_obj

import "errors"

type Username string

// Validate checks if the Username is valid
func (u *Username) Validate() error {
	if len(*u) == 0 {
		return errors.New("username cannot be empty")
	}
	if len(*u) < 3 || len(*u) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	return nil
}

// Value implements the database driver interface for Username
func (u *Username) Value() (string, error) {
	if err := u.Validate(); err != nil {
		return "", err
	}
	return string(*u), nil
}

// Scan implements the database driver interface for Username
func (u *Username) Scan(value interface{}) error {
	if value == nil {
		*u = ""
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("username must be a string")
	}

	*u = Username(str)
	return u.Validate()
}
