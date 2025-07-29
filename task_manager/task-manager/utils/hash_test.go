package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	
	hash, err := HashPassword(password)
	
	assert.NoError(t, err)
	assert.NotEqual(t, password, hash)
	assert.Len(t, hash, 60) // bcrypt hash length
}

func TestCheckPassword(t *testing.T) {
	password := "testpassword123"
	
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	
	// Test correct password
	assert.True(t, CheckPassword(password, hash))
	
	// Test incorrect password
	assert.False(t, CheckPassword("wrongpassword", hash))
}

func TestValidatePassword(t *testing.T) {
	// Test valid password
	err := ValidatePassword("password123")
	assert.NoError(t, err)
	
	// Test password too short
	err = ValidatePassword("123")
	assert.Error(t, err)
	assert.IsType(t, ErrPasswordTooShort, err)
	
	// Test empty password
	err = ValidatePassword("")
	assert.Error(t, err)
	assert.IsType(t, ErrPasswordTooShort, err)
} 