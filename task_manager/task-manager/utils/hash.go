package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword checks if a password matches its hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return ErrPasswordTooShort
	}
	return nil
}

// Common errors
var (
	ErrPasswordTooShort = &ValidationError{Message: "password must be at least 6 characters"}
)

// ValidationError represents validation errors
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
} 