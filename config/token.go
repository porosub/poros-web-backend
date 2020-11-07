package config

import (
	"os"

	"github.com/joho/godotenv"
)

// TokenENV ... TokenENV struct declaration
type TokenENV struct {
	JWTSecret string
}

// TokenENVInterface ... TokenENV interface declaration
type TokenENVInterface interface {
	BuildTokenConfig() error
}

// BuildTokenConfig ... Create token config
func (te *TokenENV) BuildTokenConfig() (*TokenENV, error) {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &TokenENV{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}
