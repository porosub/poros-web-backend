package post

import (
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/models/base"
	"github.com/divisi-developer-poros/poros-web-backend/models/post"
	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (h *PostHandler) sendSuccess(c *gin.Context, data interface{}) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, data)
}

func (h *PostHandler) sendError(c *gin.Context, status int, message string) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "error", message, status, nil)
}

// List ... Get all posts
func (h *PostHandler) List(c *gin.Context) {
	data, err := h.PostModel.List()
	if err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
	}
	for id := range *data {
		(*data)[id].LocalizedField()
	}
	h.sendSuccess(c, data)
}

// Get ... Get single post
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
	data.LocalizedField()
	h.sendSuccess(c, data)
	return
}

// Create ... Create single post
func (h *PostHandler) Create(c *gin.Context) {
	// Bind Post Data
	var p post.Post
	if err := c.ShouldBindWith(&p, binding.FormMultipart); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get User ID
	idInt, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	p.UserID = uint(idInt)

	// Get images data
	form, err := c.MultipartForm()
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	imagesBlob := form.File["images"]

	// Store images and store data to DB
	if _, err = h.PostModel.Create(&p, imagesBlob); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get newly created data
	data, err := h.PostModel.Get(p.ID)
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}

	data.LocalizedField()
	h.sendSuccess(c, data)
}

// Update ... update single post
func (h *PostHandler) Update(c *gin.Context) {
	// Get User ID
	userIDInt, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Bind Post Data
	var p post.Post
	if err := c.ShouldBindWith(&p, binding.FormMultipart); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get Body User ID
	idInt, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	p.UserID = uint(idInt)

	// Get images data
	form, err := c.MultipartForm()
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	imagesBlob := form.File["images"]

	// Update post
	if _, err = h.PostModel.Update(&p, userIDInt, imagesBlob); err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.LocalizedField()
	h.sendSuccess(c, p)
}

// Delete ... Delete single post
func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.PostModel.Delete(uint(id)); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	h.sendSuccess(c, nil)
	return
}

// AttachTags attach tags to corresponding post
func (h *PostHandler) AttachTags(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("post_id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	var t []base.Tag
	if err := c.ShouldBindJSON(&t); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err := h.PostModel.AttachTags(id, &t)
	if err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.LocalizedField()
	h.sendSuccess(c, p)
}

// DetachTags detach tags from corresponding post
func (h *PostHandler) DetachTags(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("post_id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	var t []base.Tag
	if err := c.ShouldBindJSON(&t); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err := h.PostModel.DetachTags(id, &t)
	if err != nil {
		h.sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.LocalizedField()
	h.sendSuccess(c, p)
}
