package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
)

func GetAll(c *gin.Context) {
	var users []userModel.User

	err := userModel.GetAll(&users)


	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message" : "null",
			"data": users,
 		})
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