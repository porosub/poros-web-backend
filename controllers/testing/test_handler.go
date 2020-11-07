package testing

import (
	"net/http"

	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// Cobs ... Cobs struct declaration
type Cobs struct {
	Res   response.Response
	Token token.JWTToken
}

// TestInterface ... Test interface declaration
type TestInterface interface {
	Guest(c *gin.Context)
	Login(c *gin.Context)
	Home(c *gin.Context)
}

// Guest ... Show guest menu
func (cobs *Cobs) Guest(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "guest")
	return
}

// UserTesting ... User testing struct declaration
type UserTesting struct {
	Username string `json:"username"`
	Usertype int    `json:"usertype"`
	Pass     string `json:"pass"`
}

// Login ... Login user
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

// Home ... Show home
func (cobs *Cobs) Home(c *gin.Context) {
	cobs.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, "home")
	return
}
