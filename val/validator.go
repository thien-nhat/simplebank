package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var(
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString

)
func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("string length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value,3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must only lowercase letters, numbers, and underscores")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value,3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must only letters, and spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 1000)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value,3, 100); err != nil {
		return err
	}
	if _, err:= mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidateEmailID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("email id must be greater than 0")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}