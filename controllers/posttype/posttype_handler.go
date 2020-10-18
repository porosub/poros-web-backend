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

func (h *PostTypeHandler) List(c *gin.Context) {
	data, err := h.PostTypeModel.List()
	if err != nil {
		h.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
	}
	h.Res.CustomResponse(c, "Content-Type", "application/json", "success", "", http.StatusOK, data)
}

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
		"Created.", http.StatusCreated, result)
	return
}

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
		"Updated.", http.StatusOK, result)
	return
}

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
		"Deleted.", http.StatusNoContent, nil)
	return
}
