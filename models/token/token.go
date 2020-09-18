package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/divisi-developer-poros/poros-web-backend/models/response"
	"time"
)

var ResponseEntity response.Response

func (jt *JWTToken) GenerateToken(userName string, userType int) (string, error) {
	var signingKey []byte
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTToken{
		Username: userName,
		Usertype: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "poros",
			IssuedAt:  time.Now().Unix(),
		},
	}).SignedString(signingKey)
}

func (jt *JWTToken) TokenValidation(encodedToken string) (jwt.Claims, error) {
	var signingKey []byte
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
