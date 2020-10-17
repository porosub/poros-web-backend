package token

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	Id       string `json:"id"`
	jwt.StandardClaims
}

type AccessToken struct {
	Id        string `gorm:"primaryKey"`
	ExpiresAt int64
}

type JWTTokenInterface interface {
	GenerateToken(userName string, userType int) (string, error)
	TokenValidation(token string) (*jwt.Token, error)
}
