package errors

// UserNotFoundError occurs when user is not found
type UserNotFoundError struct{}

func (e UserNotFoundError) Error() string {
	return "user not found"
}

// EmailAlreadyExistsError occurs when email is already registered
type EmailAlreadyExistsError struct{}

func (e EmailAlreadyExistsError) Error() string {
	return "email already exists"
}

// InvalidCredentialsError occurs when login credentials are invalid
type InvalidCredentialsError struct{}

func (e InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

// UserPromotionError occurs when user promotion fails
type UserPromotionError struct {
	Message string
}

func (e UserPromotionError) Error() string {
	return e.Message
} 