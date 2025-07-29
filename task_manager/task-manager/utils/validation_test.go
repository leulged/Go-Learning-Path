package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	// Test valid emails
	validEmails := []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"user+tag@example.org",
		"123@test.com",
	}

	for _, email := range validEmails {
		err := ValidateEmail(email)
		assert.NoError(t, err, "Email should be valid: %s", email)
	}

	// Test invalid emails
	invalidEmails := []string{
		"",
		"invalid-email",
		"@example.com",
		"test@",
		"test..test@example.com",
	}

	for _, email := range invalidEmails {
		err := ValidateEmail(email)
		assert.Error(t, err, "Email should be invalid: %s", email)
	}
}

func TestValidateName(t *testing.T) {
	// Test valid names
	validNames := []string{
		"John Doe",
		"Alice",
		"Bob Smith",
		"José María",
	}

	for _, name := range validNames {
		err := ValidateName(name)
		assert.NoError(t, err, "Name should be valid: %s", name)
	}

	// Test invalid names
	invalidNames := []string{
		"",
		"A", // too short
		"Very long name that exceeds the maximum allowed length of fifty characters", // too long
	}

	for _, name := range invalidNames {
		err := ValidateName(name)
		assert.Error(t, err, "Name should be invalid: %s", name)
	}
}

func TestValidateTaskTitle(t *testing.T) {
	// Test valid titles
	validTitles := []string{
		"Complete project",
		"Review code",
		"Update documentation",
	}

	for _, title := range validTitles {
		err := ValidateTaskTitle(title)
		assert.NoError(t, err, "Title should be valid: %s", title)
	}

	// Test invalid titles
	invalidTitles := []string{
		"",
		"Very long task title that exceeds the maximum allowed length of one hundred characters and should be rejected by the validation function",
	}

	for _, title := range invalidTitles {
		err := ValidateTaskTitle(title)
		assert.Error(t, err, "Title should be invalid: %s", title)
	}
}

func TestValidateTaskStatus(t *testing.T) {
	// Test valid statuses
	validStatuses := []string{
		"Pending",
		"In Progress",
		"Completed",
	}

	for _, status := range validStatuses {
		err := ValidateTaskStatus(status)
		assert.NoError(t, err, "Status should be valid: %s", status)
	}

	// Test invalid statuses
	invalidStatuses := []string{
		"pending",
		"in progress",
		"completed",
		"Done",
		"",
	}

	for _, status := range invalidStatuses {
		err := ValidateTaskStatus(status)
		assert.Error(t, err, "Status should be invalid: %s", status)
	}
} 