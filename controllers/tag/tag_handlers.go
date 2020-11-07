package tag

import (
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/models/tags"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// HandlerTag ... Tag handler struct declaration
type HandlerTag struct {
	TagModel tags.Tag
	Res      response.Response
}

// HandlerTagInterface ... Tag handler interface declaration
type HandlerTagInterface interface {
	GetTags(c *gin.Context)
	GetTagByID(c *gin.Context)
	CreateTag(c *gin.Context)
	UpdateTagByID(c *gin.Context)
	DeleteTag(c *gin.Context)
}

// GetTags ... Get all tags
func (ht *HandlerTag) GetTags(c *gin.Context) {
	data, err := ht.TagModel.FetchTags()
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusOK, nil)
		return
	}
	ht.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, data)
	return
}

// GetTagByID ... Get single tag
func (ht *HandlerTag) GetTagByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusBadRequest, nil)
		return
	}

	data, err := ht.TagModel.FetchTagByID(id)
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusNotFound, nil)
		return
	}

	ht.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, data)
	return
}

// CreateTag ... Create single tag
func (ht *HandlerTag) CreateTag(c *gin.Context) {
	var newTag tags.Tag
	if err := c.ShouldBindJSON(&newTag); err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusBadRequest, nil)
		return
	}

	result, err := ht.TagModel.CreateTag(&newTag)
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusInternalServerError, nil)
		return
	}

	ht.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, result)
	return
}

// UpdateTagByID ... Update single tag
func (ht *HandlerTag) UpdateTagByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusBadRequest, nil)
		return
	}

	var updatedTag tags.Tag
	if err := c.ShouldBindJSON(&updatedTag); err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusBadRequest, nil)
		return
	}
	updatedTag.ID = id

	result, err := ht.TagModel.UpdateTagByID(&updatedTag)
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusInternalServerError, nil)
		return
	}

	ht.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, result)
	return
}

// DeleteTag ... Delete single tag
func (ht *HandlerTag) DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			"Bad Request Params", http.StatusBadRequest, nil)
		return
	}

	if err := ht.TagModel.DeleteTagByID(id); err != nil {
		ht.Res.CustomResponse(c, "Content-Type",
			"application/json", "error",
			err.Error(), http.StatusNotFound, nil)
		return
	}

	ht.Res.CustomResponse(c, "Content-Type",
		"application/json", "success",
		"", http.StatusOK, nil)
	return
}
