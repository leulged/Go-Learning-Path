package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return ErrEmailRequired
	}

	// Check for consecutive dots in local part
	if strings.Contains(email, "..") {
		return ErrInvalidEmail
	}

	// Basic email regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

// ValidateName validates user name
func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrNameRequired
	}

	if len(name) < 2 {
		return ErrNameTooShort
	}

	if len(name) > 50 {
		return ErrNameTooLong
	}

	return nil
}

// ValidateTaskTitle validates task title
func ValidateTaskTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrTitleRequired
	}

	if len(title) > 100 {
		return ErrTitleTooLong
	}

	return nil
}

// ValidateTaskStatus validates task status
func ValidateTaskStatus(status string) error {
	validStatuses := []string{"Pending", "In Progress", "Completed"}
	
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}
	
	return ErrInvalidStatus
}

// Common validation errors
var (
	ErrEmailRequired   = &ValidationError{Message: "email is required"}
	ErrInvalidEmail    = &ValidationError{Message: "invalid email format"}
	ErrNameRequired    = &ValidationError{Message: "name is required"}
	ErrNameTooShort    = &ValidationError{Message: "name must be at least 2 characters"}
	ErrNameTooLong     = &ValidationError{Message: "name must be less than 50 characters"}
	ErrTitleRequired   = &ValidationError{Message: "title is required"}
	ErrTitleTooLong    = &ValidationError{Message: "title must be less than 100 characters"}
	ErrInvalidStatus   = &ValidationError{Message: "invalid status"}
) 