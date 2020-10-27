package post

import (
	"net/http"
	"strconv"

	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/models/post"
	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
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
	var p post.Post
	if err := c.ShouldBindWith(&p, binding.FormMultipart); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	idInt, _ := strconv.Atoi(c.PostForm("user_id"))
	if err := user.Get(&user.User{}, idInt); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	p.UserID = uint(idInt)

	filenames, err := storage.StoreFiles(c, "images", config.AssetPostsImages)
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	h.PostModel.LinkImagesName(&p, *filenames)

	if _, err = h.PostModel.Create(&p); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	h.sendSuccess(c, "", p)
}

func (h *PostHandler) Update(c *gin.Context) {
	// Mengambil post lama
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err := h.PostModel.Get(uint(id))
	if err != nil {
		h.sendError(c, http.StatusNotFound, err.Error())
		return
	}

	// Mengambil post baru
	if err := c.ShouldBindWith(&p, binding.FormMultipart); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	idInt, _ := strconv.Atoi(c.PostForm("user_id"))
	if err := user.Get(&user.User{}, idInt); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	p.UserID = uint(idInt)

	// Hapus gambar lama
	if err := h.PostModel.DeletePostImages(p); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Menyimpan gambar baru
	filenames, err := storage.StoreFiles(c, "images", config.AssetPostsImages)
	if err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	h.PostModel.LinkImagesName(p, *filenames)

	// Mengubah post di db
	if _, err := h.PostModel.Update(p); err != nil {
		h.sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	h.sendSuccess(c, "", p)
}

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

	h.sendSuccess(c, "", nil)
	return
}
