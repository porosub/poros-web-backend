package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/Hash"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	Token token.JWTToken
	Res   response.Response
}

type AuthHandlersInterface interface {
	Login(c *gin.Context)
	Me(c *gin.Context)
	Logout(c *gin.Context)
}

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

func (a *AuthHandlers) Login(c *gin.Context) {
	var data LoginForm
	if err := c.ShouldBind(&data); err != nil {
		a.sendError(c, 422, err.Error())
		return
	}

	var usr user.User
	if err := user.GetByUsername(&usr, data.Username); err != nil {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	if Hash.GetSha1Hash(data.Password) != usr.Password {
		a.sendError(c, 401, Hash.GetSha1Hash(data.Password))
		return
	}
	accessToken, err := a.Token.GenerateToken(data.Username, usr.User_type_id)
	if err != nil {
		a.sendError(c, 500, err.Error())
		return
	}
	a.sendSuccess(c, "", gin.H{
		"access_token": accessToken,
	})
}

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
	tokenId, ok := claims["id"].(string)
	if !ok {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	if err := jwtModel.DeleteToken(tokenId); err != nil {
		a.sendError(c, 401, "Unauthorized.")
		return
	}
	a.sendSuccess(c, "Logged out.", nil)
}

func (a *AuthHandlers) Me(c *gin.Context) {
	var authUser user.User
	if err := user.GetByUsername(&authUser, c.Writer.Header().Get("User")); err != nil {
		a.sendError(c, 500, err.Error())
	}
	authUser.Password = ""
	a.sendSuccess(c, "", &authUser)
}
