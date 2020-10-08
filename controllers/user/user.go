package user

import (
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Model	userModel.User
	Res     r.Response
}

func (usr *UserHandler) GetAll(c *gin.Context) {
	var users []userModel.User

	if err := userModel.GetAll(&users); err != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
	} else {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, users)
		return
	}

}

func (usr *UserHandler) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var user userModel.User
		if err := userModel.Get(&user, numId); err != nil {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
		} else {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, user)
			return
		}
	}
}

func (usr *UserHandler) Create(c *gin.Context) {
	var user userModel.User

	if err := c.ShouldBindJSON(&user); err != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
	} else {
		if image, err := c.FormFile("image"); err != nil {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "no file provided", http.StatusBadRequest, nil)
		} else {
			imageUrl := "images/" + image.Filename
			if err := c.SaveUploadedFile(image, imageUrl); err != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "error when uploading images", http.StatusInternalServerError, nil)
			} else {
				user.Image = imageUrl
				if err := userModel.Create(&user); err != nil {
					usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
				} else {
					usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user created", http.StatusOK, user)
					return
				}
			}
		}
	}
}

func (usr *UserHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var user userModel.User

	if err := c.ShouldBindJSON(&user); err != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
	} else {
		if numId, error := strconv.Atoi(id); error != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
		} else {
			var existedUser userModel.User
			if errUserExist := userModel.Get(&existedUser, numId); errUserExist != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
				return
			}
			if err := userModel.Update(&user, numId); err != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusBadRequest, nil)
				return
			} else {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user updated", http.StatusOK, user)
				return
			}
		}
	}
}

func (usr *UserHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var user userModel.User

		if errUserExist := userModel.Get(&user, numId); errUserExist != nil {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
		} else {
			if err := userModel.Delete(&user, numId); err != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
			} else {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user deleted", http.StatusOK, nil)
				return
			}
		}
	}
}
