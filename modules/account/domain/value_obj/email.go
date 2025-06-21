package value_obj

import "errors"

type Email string

// Validate checks if the Email is valid
func (e *Email) Validate() error {
	if len(*e) == 0 {
		return errors.New("email cannot be empty")
	}
	if len(*e) < 5 || len(*e) > 254 {
		return errors.New("email must be between 5 and 254 characters")
	}

	// TODO: Add more complex email validation logic here if needed

	return nil
}

// Value implements the database driver interface for Email
func (e *Email) Value() (string, error) {
	if err := e.Validate(); err != nil {
		return "", err
	}
	return string(*e), nil
}

// Scan implements the database driver interface for Email
func (e *Email) Scan(value interface{}) error {
	if value == nil {
		*e = ""
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("email must be a string")
	}

	*e = Email(str)
	return e.Validate()
}
