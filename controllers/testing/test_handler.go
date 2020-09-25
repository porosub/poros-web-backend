package testing

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cobs struct {
	Res     response.Response
	Token   token.JWTToken
}

type TestInterface interface {
	Guest(c *gin.Context)
	Login(c *gin.Context)
	Home(c *gin.Context)
}

func (cobs *Cobs) Guest(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "guest")
	return
}

type UserTesting struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	Pass     string `json:"pass"`
}

func (cobs *Cobs) Login(c *gin.Context) {
	var userCobs UserTesting
	if err := c.ShouldBindJSON(c); err != nil {
		cobs.Res.CustomResponse(c, "Content-Type", "application/json", "error", "failed when binding data", http.StatusBadRequest, nil)
		return
	}

	value, err := cobs.Token.GenerateToken(userCobs.Username, userCobs.Usertype)
	if err != nil {
		cobs.Res.CustomResponse(c, "Content-Type", "application/json", "error", "failed when generating token", http.StatusInternalServerError, nil)
		return
	}
	cobs.Res.CustomResponse(c, "Authorization", "Bearer "+value, "success", "", http.StatusOK, value)
	return
}

func (cobs *Cobs) Home(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "home")
	return
}
