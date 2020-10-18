package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/models/post"
	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostModel     post.Post
	PostTypeModel posttype.PostType
	Res           response.Response
}

type PostHandlerInterface interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func (h *PostHandler) sendSuccess(c *gin.Context, message string, data interface{}) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", message, http.StatusOK, data)
}

func (h *PostHandler) sendCreated(c *gin.Context, message string, data interface{}) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", message, http.StatusCreated, data)
}

func (h *PostHandler) sendError(c *gin.Context, status int, message string) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "error", message, status, nil)
}

func (h *PostHandler) List(c *gin.Context) {
	data, err := h.PostModel.List()
	if err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
	}
	h.sendSuccess(c, "", data)
}

func (h *PostHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := h.PostModel.Get(uint(id))
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}
	h.sendSuccess(c, "", data)
	return
}

func (h *PostHandler) Create(c *gin.Context) {
	var data post.Post
	if err := c.ShouldBind(&data); err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Assign authenticated user to post creator
	var usr user.User
	if err := user.GetByUsername(&usr, c.Writer.Header().Get("User")); err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	data.UserID = usr.Id
	usr.Password = ""
	data.User = usr
	result, err := h.PostModel.Create(&data)

	// Validate post type
	postType, err := h.PostTypeModel.Get(data.PostTypeID)
	if err != nil {
		h.sendError(c, http.StatusBadRequest, fmt.Sprintf("Post Type ID %d not found.", data.PostTypeID))
		return
	}
	data.PostType = *postType

	h.sendSuccess(c, "Created", result)
	return
}

func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var data post.Post
	if err := c.ShouldBind(&data); err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	data.ID = uint(id)

	postData, err := h.PostModel.Get(uint(id))
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}
	data.CreatedAt = postData.CreatedAt

	// Validate user
	var usr user.User
	if err := user.GetByUsername(&usr, c.Writer.Header().Get("User")); err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if postData.UserID != usr.Id {
		h.sendError(c, http.StatusForbidden, "This post isn't your owns")
		return
	}
	data.UserID = usr.Id
	data.User = usr

	// Validate post type
	postType, err := h.PostTypeModel.Get(data.PostTypeID)
	if err != nil {
		h.sendError(c, http.StatusBadRequest, fmt.Sprintf("Post Type ID %d not found.", data.PostTypeID))
		return
	}
	data.PostType = *postType

	result, err := h.PostModel.Update(&data)
	if err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	result.User.Password = ""
	result.PostImage = postData.PostImage
	h.sendSuccess(c, "Updated", result)
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get post
	postData, err := h.PostModel.Get(uint(id))
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}

	if postData.User.Username != c.Writer.Header().Get("User") {
		h.sendError(c, http.StatusForbidden, "This post isn't your own.")
		return
	}

	if err := h.PostModel.Delete(uint(id)); err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"Deleted.", http.StatusNoContent, nil)
	return
}
