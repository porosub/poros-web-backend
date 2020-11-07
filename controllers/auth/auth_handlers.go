package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/host"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// AuthHandlers ... Struct for Auth Handler
type AuthHandlers struct {
	Token token.JWTToken
	Res   response.Response
}

// AuthHandlersInterface ... Interface for Auth Handler
type AuthHandlersInterface interface {
	Login(c *gin.Context)
	Me(c *gin.Context)
	Logout(c *gin.Context)
}

// LoginForm ... Declare struct for storing request
type LoginForm struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (a *AuthHandlers) sendSuccess(c *gin.Context, message string, data interface{}) {
	a.Res.CustomResponse(c, "Content-Type", "application/json", "success", message, http.StatusOK, data)
}

func (a *AuthHandlers) sendError(c *gin.Context, status int, message string) {
	a.Res.CustomResponse(c, "Content-Type", "application/json", "error", message, status, nil)
}

// Login ... user login
func (a *AuthHandlers) Login(c *gin.Context) {
	fmt.Printf("Host: %v\n", host.Host)
	var loginForm LoginForm
	var u user.User
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		a.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.SignIn(&u, loginForm.Username, loginForm.Password); err != nil {
		a.sendError(c, http.StatusNotFound, err.Error())
		return
	}

	accessToken, err := a.Token.GenerateToken(u.Username, u.UserTypeID)
	if err != nil {
		a.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.Password = ""
	u.LocalizedField()
	a.sendSuccess(c, "", gin.H{
		"user":         u,
		"access_token": accessToken,
	})

}

// Logout ... user logout
func (a *AuthHandlers) Logout(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if len(accessToken) == 0 {
		a.sendError(c, 401, "missing authorization header")
		return
	}
	accessToken = strings.Replace(accessToken, "Bearer ", "", 1)
	jwtModel := token.JWTToken{}
	result, err := jwtModel.TokenValidation(accessToken)
	if err != nil {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	claims, ok := result.Claims.(jwt.MapClaims)
	if !ok {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	tokenID, ok := claims["id"].(string)
	if !ok {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	if err := jwtModel.DeleteToken(tokenID); err != nil {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	a.sendSuccess(c, "Logged out.", nil)
}

// Me ... Show User Data
func (a *AuthHandlers) Me(c *gin.Context) {
	var authUser user.User
	if err := user.GetByUsername(&authUser, c.Writer.Header().Get("User")); err != nil {
		a.sendError(c, 500, err.Error())
	}
	authUser.Password = ""
	a.sendSuccess(c, "", &authUser)
}
