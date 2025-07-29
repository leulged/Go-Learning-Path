package entities

// User is the core domain entity - pure business logic
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     string
}

// NewUser creates a new user with validation
func NewUser(name, email, password string) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "user", // default role
	}
}

// IsAdmin checks if user has admin role
func (u User) IsAdmin() bool {
	return u.Role == "admin"
}

// SetRole sets the user role
func (u *User) SetRole(role string) {
	u.Role = role
} 