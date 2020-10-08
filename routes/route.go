package routes

import (
	"net/http"
	"github.com/divisi-developer-poros/poros-web-backend/controllers/tag"
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	userController "github.com/divisi-developer-poros/poros-web-backend/controllers/user"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"
	"github.com/gin-gonic/gin"

)

var (
	TestingHandlers test.Cobs
	TagHandlers     tag.HandlerTag
	TokenMiddleware middleware.TokenMiddleware
)

type Test struct {
	message	string
	status int
}

// Start inisialisasi route yang digunakan
func Start() {
	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})


	// Example Of JWT Middleware
	r.GET("/guest", TestingHandlers.Guest)
	r.POST("/login", TestingHandlers.Login)
	r.GET("/home", TokenMiddleware.AuthorizeToken, TestingHandlers.Home)

	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.Get)
	r.POST("/users", userController.Create)
	r.PUT("/users/:id", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	// Example of Upload File
	r.POST("/upload-file", storage.SingleHandler)
	// Example of Upload Files
	r.POST("/upload-files", storage.MultipleHandler)
	// Static files for assets
	r.Static("/assets", "./assets")

	// tag routes
	r.GET("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.GetTags)
	r.GET("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.GetTagByID)
	r.POST("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.CreateTag)
	r.PUT("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.UpdateTagByID)
	r.DELETE("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.DeleteTag)

	r.Run()
}
