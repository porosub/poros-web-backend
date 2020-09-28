package user

import (
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"net/http"
	"strconv"
)

type Response struct {
	Res     r.Response
}

func (response *Response) GetAll(c *gin.Context) {
	var users []userModel.User

	if err := userModel.GetAll(&users); err != nil {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
	} else {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, users)
		return
	}

}

func (response *Response) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var user userModel.User
		if err := userModel.Get(&user, numId); err != nil {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
		} else {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, user)
			return
		}
	}
}

func (response *Response) Create(c *gin.Context) {
	var user userModel.User

	c.BindJSON(&user)

	validate := validator.New()

	if errValidate:= validate.Struct(user); errValidate != nil {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "invalid user input", http.StatusBadRequest, nil)
	} else {
		if err := userModel.Create(&user); err != nil {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
		} else {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user created", http.StatusOK, user)
			return
		}
	}
}

func (response *Response) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var user userModel.User
	c.BindJSON(&user)
	validate := validator.New()

	if errValidate:= validate.Struct(user); errValidate != nil {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "invalid user input", http.StatusBadRequest, nil)
	} else {
		if numId, error := strconv.Atoi(id); error != nil {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
		} else {
			var existedUser userModel.User
			if errUserExist := userModel.Get(&existedUser, numId); errUserExist != nil {
				response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
				return
			}
			if err := userModel.Update(&user, numId); err != nil {
				response.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
				return
			} else {
				response.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user updated", http.StatusOK, user)
				return
			}
		}
	}
}

func (response *Response) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var user userModel.User

		if errUserExist := userModel.Get(&user, numId); errUserExist != nil {
			response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
		} else {
			if err := userModel.Delete(&user, numId); err != nil {
				response.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
			} else {
				response.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user deleted", http.StatusOK, nil)
				return
			}
		}
	}
}
