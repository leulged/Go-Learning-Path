package interfaces

// TokenService interface defines JWT token operations
type TokenService interface {
	GenerateToken(email, role string) (string, error)
	ValidateToken(token string) (string, string, error) // returns email, role, error
	ExtractClaims(token string) (map[string]interface{}, error)
} 