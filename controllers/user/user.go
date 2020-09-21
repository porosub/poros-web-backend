package user

import (
	"fmt"
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	var users []userModel.User

	err := userModel.GetAll(&users)

	if err != nil {
		response.Res(c, "Content-Type", "application/json", "error", "failed when fetching users", http.StatusBadRequest, err)
		return
	} else {
		response.Res(c, "Content-Type", "application/json", "success", "null", http.StatusOK, users)
		return
	}

}

func Get(c *gin.Context) {
	id := c.Params.ByName("id")

	_, error := strconv.Atoi(id)

	if error != nil {
		response.Res(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
		return
	} else {
		var user userModel.User

		err := userModel.Get(&user, id)

		if err != nil {
			response.Res(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, err)
			return
		} else {
			response.Res(c, "Content-Type", "application/json", "success", "null", http.StatusOK, user)
			return
		}
	}
}

func Create(c *gin.Context) {
	var user userModel.User

	c.BindJSON(&user)

	fmt.Println(user)

	return
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}