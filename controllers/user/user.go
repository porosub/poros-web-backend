package user

import (
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	var users []userModel.User

	err := userModel.GetAll(&users)

	if err != nil {
		response.Res(c, "Content-Type", "application/json", "error", "failed when fetching users", http.StatusBadRequest, err)
	} else {
		response.Res(c, "Content-Type", "application/json", "success", "null", http.StatusOK, users)
		return
	}

}

func Get(c *gin.Context) {

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}