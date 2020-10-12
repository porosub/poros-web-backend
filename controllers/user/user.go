package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/random"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UserHandler struct {
	Model	userModel.User
	Res     r.Response
}

func (usr *UserHandler) GetAll(c *gin.Context) {
	var users []userModel.User

	if err := userModel.GetAll(&users); err != nil {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
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
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
		} else {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, user)
			return
		}
	}
}

func (usr *UserHandler) Create(c *gin.Context) {
	if isOk, user := _checkUserBinding(c); isOk != true {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "error when binding user", http.StatusBadRequest, nil)
	} else {
		if image, errImg := c.FormFile("image"); errImg != nil {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", errImg.Error(), http.StatusBadRequest, nil)
		} else {
			contentType := strings.Split(image.Header.Get("Content-Type"), "/")
			filename := random.RandomString(32) + "." + contentType[1]
			imageUrl := config.AssetUsersImages + filename
			if err := c.SaveUploadedFile(image, imageUrl); err != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
			} else {
				user.Image = imageUrl
				if err := userModel.Create(&user); err != nil {
					usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
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

	if isOk, user := _checkUserBinding(c); isOk != true {
		usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "error when binding user", http.StatusBadRequest, nil)
	} else {
		if numId, error := strconv.Atoi(id); error != nil {
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
		} else {
			var existedUser userModel.User
			if errUserExist := userModel.Get(&existedUser, numId); errUserExist != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", errUserExist.Error(), http.StatusInternalServerError, nil)
			} else {
				if image, err := c.FormFile("image"); err != nil {
					user.Image = existedUser.Image
				} else {
					contentType := strings.Split(image.Header.Get("Content-Type"), "/")
					filename := random.RandomString(32) + "." + contentType[1]
					imageUrl := config.AssetUsersImages + filename
					if err := c.SaveUploadedFile(image, imageUrl); err != nil {
						usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
					} else {
						user.Image = imageUrl
						os.Remove(existedUser.Image)
					}
				}
				if err := userModel.Update(&user, numId); err != nil {
					usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
				} else {
					user.Id = existedUser.Id
					usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user updated", http.StatusOK, user)
				}
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
			usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", errUserExist.Error(), http.StatusInternalServerError, nil)
		} else {
			userImageUrl := user.Image
			if err := userModel.Delete(&user, numId); err != nil {
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
			} else {
				os.Remove(userImageUrl)
				usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user deleted", http.StatusOK, nil)
				return
			}
		}
	}
}

func _checkUserBinding(c *gin.Context) (bool bool, user userModel.User) {
	if len(c.PostForm("username")) <= 0 || len(c.PostForm("password")) <= 0 || len(c.PostForm("full_name")) <= 0 || len(c.PostForm("user_type_id")) <= 0 {
		return false, user
	}
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Full_name = c.PostForm("full_name")
	user.User_type_id, _ = strconv.Atoi(c.PostForm("user_type_id"))
	return true, user
}
