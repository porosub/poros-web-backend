package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jt "github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// TokenMiddleware ... Token middleware struct declaration
type TokenMiddleware struct {
	ResponseEntity response.Response
	JWT            jt.JWTToken
}

// AuthorizeToken ... Token Validation
func (tm *TokenMiddleware) AuthorizeToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		tm.ResponseEntity.CustomResponse(c,
			"Content-Type",
			"application/json",
			"error",
			"missing authorization header",
			http.StatusUnauthorized,
			nil)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := tm.JWT.TokenValidation(token)
	if err != nil {
		tm.ResponseEntity.CustomResponse(
			c,
			"Content-Type",
			"application/json",
			"error",
			"failed when validating token",
			http.StatusUnauthorized,
			nil)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	clms, ok := claims.Claims.(jwt.MapClaims)
	if ok {
		username, ok := clms["username"].(string)
		if ok {
			c.Writer.Header().Set("User", username)
		}
	}
}
