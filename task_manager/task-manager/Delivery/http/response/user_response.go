package response

import "task_manager/Domain/entities"

// UserResponse represents the user data sent in HTTP responses
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// ToUserResponse converts domain User to UserResponse
func ToUserResponse(user entities.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

// LoginResponse represents the login response
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

// ToLoginResponse creates a login response
func ToLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Success: true,
		Token:   token,
	}
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}

// ToMessageResponse creates a message response
func ToMessageResponse(message string) MessageResponse {
	return MessageResponse{
		Message: message,
	}
} 