package postimage

import (
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/models/postimage"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// PostImageHandler ... Post image handler struct declaration
type PostImageHandler struct {
	PostImageModel postimage.PostImage
	Res            response.Response
}

// PostImageHandlerInterface ... Post image handler struct interface
type PostImageHandlerInterface interface {
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func (h *PostImageHandler) sendSuccess(c *gin.Context, message string, data interface{}) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", message, http.StatusOK, data)
}

func (h *PostImageHandler) sendError(c *gin.Context, status int, message string) {
	h.Res.CustomResponse(c, "Content-Type", "application/json", "error", message, status, nil)
}

// Get ... Get single post image
func (h *PostImageHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := h.PostImageModel.Get(uint(id))
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}

	h.sendSuccess(c, "", data)
}

// Update ... Update single post image
func (h *PostImageHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}

	var data postimage.PostImage
	if err := c.ShouldBindJSON(&data); err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}
	data.ID = uint(id)

	result, err := h.PostImageModel.Update(&data)
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"Updated.", http.StatusOK, result)
	return
}

// Delete ... delete post image
func (h *PostImageHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			"Bad Request Params", http.StatusUnprocessableEntity, nil)
		return
	}

	if err := h.PostImageModel.Delete(uint(id)); err != nil {
		h.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusNotFound, nil)
		return
	}

	h.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"Deleted.", http.StatusNoContent, nil)
	return
}
