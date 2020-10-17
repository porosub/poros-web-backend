package routes

import (
	"net/http"

	"github.com/divisi-developer-poros/poros-web-backend/controllers/auth"
	"github.com/divisi-developer-poros/poros-web-backend/controllers/tag"
	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"
	userController "github.com/divisi-developer-poros/poros-web-backend/controllers/user"
	UserTypeController "github.com/divisi-developer-poros/poros-web-backend/controllers/user_type"
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
	"github.com/gin-gonic/gin"
)

var (
	TestingHandlers  test.Cobs
	TagHandlers      tag.HandlerTag
	TokenMiddleware  middleware.TokenMiddleware
	AuthHandlers     auth.AuthHandlers
	UserHandlers     userController.UserHandler
	UserTypeHandlers UserTypeController.UserTypeHandler
)

type Test struct {
	message string
	status  int
}

// Start inisialisasi route yang digunakan
func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// Example Of JWT Middleware
	r.GET("/guest", TestingHandlers.Guest)
	r.GET("/home", TokenMiddleware.AuthorizeToken, TestingHandlers.Home)

	// Auth routes
	r.POST("/auth/login", AuthHandlers.Login)
	r.GET("/auth/me", TokenMiddleware.AuthorizeToken, AuthHandlers.Me)
	r.GET("/auth/logout", AuthHandlers.Logout)

	// user routes
	r.GET("/users", TokenMiddleware.AuthorizeToken, UserHandlers.GetAll)
	r.GET("/users/:id", TokenMiddleware.AuthorizeToken, UserHandlers.Get)
	r.POST("/users", UserHandlers.Create)
	r.PUT("/users/:id", TokenMiddleware.AuthorizeToken, UserHandlers.Update)
	r.DELETE("/users/:id", TokenMiddleware.AuthorizeToken, UserHandlers.Delete)

	// user_type routes
	r.GET("/usertype", UserTypeHandlers.GetAll)
	r.GET("/usertype/:id", UserTypeHandlers.Get)
	r.POST("/usertype", UserTypeHandlers.Create)
	r.PUT("/usertype/:id", UserTypeHandlers.Update)
	r.DELETE("/usertype/:id", UserTypeHandlers.Delete)

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
