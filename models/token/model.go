package token

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	jwt.StandardClaims
}

type JWTTokenInterface interface {
	GenerateToken(userName string, userType int) (string, error)
	TokenValidation(token string) (jwt.Claims, error)
}
