package posttype

import (
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

type PostTypeHandler struct {
	PostTypeModel posttype.PostType
	Res           response.Response
}

type PostTypeHandlerInterface interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// List ... get all post type
func (h *PostTypeHandler) List(c *gin.Context) {
	data, err := h.PostTypeModel.List()
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
	}
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, data)
}

// Get ... get post type
func (h *PostTypeHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}

	data, err := h.PostTypeModel.Get(uint(id))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusNotFound, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, data)
	return
}

// Create ... create post type
func (h *PostTypeHandler) Create(c *gin.Context) {
	var data posttype.PostType
	if err := c.ShouldBind(&data); err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}

	result, err := h.PostTypeModel.Create(&data)
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusCreated, result)
	return
}

// Update ... update post type
func (h *PostTypeHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}

	var data posttype.PostType
	if err := c.ShouldBindJSON(&data); err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}
	data.ID = uint(id)

	result, err := h.PostTypeModel.Update(&data)
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, result)
	return
}

// Delete ... delete post type
func (h *PostTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			"Bad Request Params", http.StatusUnprocessableEntity, nil)
		return
	}

	if err := h.PostTypeModel.Delete(uint(id)); err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusNotFound, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, nil)
	return
}
