package token

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTToken ... JWTToken model declaration
type JWTToken struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

// AccessToken ... AccessToken model declaration
type AccessToken struct {
	ID        string `gorm:"primaryKey"`
	ExpiresAt int64
}

// JWTTokenInterface ... Interface for JWTTokenInterface
type JWTTokenInterface interface {
	GenerateToken(userName string, userType int) (string, error)
	TokenValidation(token string) (*jwt.Token, error)
}
