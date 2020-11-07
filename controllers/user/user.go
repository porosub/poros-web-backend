package user

import (
	"net/http"
	"strconv"

	userModel "github.com/divisi-developer-poros/poros-web-backend/models/user"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UserHandler ... Struct for User Handler
type UserHandler struct {
	Model userModel.User
	Res   r.Response
}

// GetAll ... Get all users
func (usr *UserHandler) GetAll(c *gin.Context) {
	var users []userModel.User

	if err := userModel.GetAll(&users); err != nil {
		usr.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	usr.responseSuccess(c, users)
}

// Get ... Get single user
func (usr *UserHandler) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		usr.responseError(c, http.StatusBadRequest, "ID not valid!")
		return
	}

	var user userModel.User
	if err := userModel.Get(&user, numID); err != nil {
		usr.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	usr.responseSuccess(c, user)
}

// Create ... Create single user
func (usr *UserHandler) Create(c *gin.Context) {
	// Bind User Data
	var u userModel.User
	if err := c.ShouldBindWith(&u, binding.FormMultipart); err != nil {
		usr.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get image data
	imageBlob, err := c.FormFile("image_blob")
	if err != nil {
		usr.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Store image and store data to DB
	if err := userModel.Create(&u, imageBlob); err != nil {
		usr.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	usr.responseSuccess(c, u)
}

// Update ... Update single user
func (usr *UserHandler) Update(c *gin.Context) {
	// Get User ID
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		usr.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Bind User Data
	var u userModel.User
	if err := c.ShouldBindWith(&u, binding.FormMultipart); err != nil {
		usr.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get image data
	imageBlob, _ := c.FormFile("image_blob")

	if err = userModel.Update(&u, idInt, imageBlob); err != nil {
		usr.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	usr.responseSuccess(c, u)
}

// Delete ... Delete single user
func (usr *UserHandler) Delete(c *gin.Context) {
	// Get User ID
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		usr.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := userModel.Delete(idInt); err != nil {
		usr.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	usr.responseSuccess(c, "")
}

func (usr *UserHandler) responseSuccess(c *gin.Context, data interface{}) {
	usr.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, data)
}

func (usr *UserHandler) responseError(c *gin.Context, code int, message string) {
	usr.Res.CustomResponse(c, "Content-Type", "application/json", "error", message, code, nil)
}
