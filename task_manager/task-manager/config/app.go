package config

import "strconv"

// AppConfig holds application configuration
type AppConfig struct {
	Port      string
	JWTSecret string
	Environment string
}

// NewAppConfig creates a new application configuration
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "your_jwt_secret_key"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// IsDevelopment checks if running in development mode
func (c *AppConfig) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction checks if running in production mode
func (c *AppConfig) IsProduction() bool {
	return c.Environment == "production"
}

// GetPort returns the port as integer
func (c *AppConfig) GetPort() int {
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		return 8080
	}
	return port
} 