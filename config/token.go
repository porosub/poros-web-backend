package config

import (
	"github.com/joho/godotenv"
	"os"
)

type TokenENV struct {
	JWTSecret string
}

type TokenENVInterface interface {
	BuildTokenConfig() error
}

func (te *TokenENV) BuildTokenConfig() (*TokenENV, error) {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &TokenENV{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}
