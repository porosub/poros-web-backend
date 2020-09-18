package testing

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/response"
	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cobs struct {
	Res   response.Response
	Token token.JWTToken
}

type TestInterface interface {
	Guest(c *gin.Context)
	Login(c *gin.Context)
	Home(c *gin.Context)
}

func (cobs *Cobs) Guest(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "guest")
}

type UserTesting struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	Pass     string `json:"pass"`
}

func (cobs *Cobs) Login(c *gin.Context) {
	var userCobs UserTesting
	if err := c.BindJSON(&userCobs); err != nil {
		cobs.Res.CustomResponse(c, "Content-Type", "application/json", "failed", "error when parsing data", http.StatusBadRequest, nil)
	}

	value, err := cobs.Token.GenerateToken(userCobs.Username, userCobs.Usertype)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	cobs.Res.CustomResponse(c, "Authorization", "Bearer "+value, "success", "", http.StatusOK, value)
}

func (cobs *Cobs) Home(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "home")
}
