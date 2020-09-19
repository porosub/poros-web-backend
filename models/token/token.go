package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func (jt *JWTToken) GenerateToken(userName string, userType int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTToken{
		Username: userName,
		Usertype: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "poros",
			IssuedAt:  time.Now().Unix(),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (jt *JWTToken) TokenValidation(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		//return signingKey, nil
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, errors.New("invalid token")
		} else {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}
	})
}
