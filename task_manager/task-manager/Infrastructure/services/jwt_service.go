package services

import (
	"time"
	"task_manager/Domain/interfaces"
	"task_manager/config"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secret []byte
}

// NewJWTService creates a new JWT service with secret from config
func NewJWTService() interfaces.TokenService {
	appConfig := config.NewAppConfig()
	
	return &jwtService{
		secret: []byte(appConfig.JWTSecret),
	}
}

type CustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *jwtService) ValidateToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})

	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", "", jwt.ErrSignatureInvalid
	}

	return claims.Email, claims.Role, nil
}

func (j *jwtService) ExtractClaims(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return map[string]interface{}{
		"email": claims.Email,
		"role":  claims.Role,
		"exp":   claims.ExpiresAt,
	}, nil
} 