package middleware

import (
	"fmt"
	"github.com/divisi-developer-poros/poros-web-backend/models/response"
	jt "github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TokenMiddleware struct {
	ResponseEntity response.Response
	JWT jt.JWTToken
}

func (tm *TokenMiddleware) AuthorizeToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		tm.ResponseEntity.CustomResponse(c,
			"Content-Type",
			"application/json",
			"failed",
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
			"failed",
			"failed when validating token",
			http.StatusUnauthorized,
			nil)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		fmt.Println(claims)
	}
}
