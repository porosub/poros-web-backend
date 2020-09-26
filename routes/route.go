package routes

import (
	"github.com/divisi-developer-poros/poros-web-backend/controllers/tag"
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	"net/http"

	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"
	"github.com/gin-gonic/gin"
)

var (
	TestingHandlers test.Cobs
	TagHandlers     tag.HandlerTag
	TokenMiddleware middleware.TokenMiddleware
)

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

	// tag routes
	r.GET("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.GetTags)
	r.GET("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.GetTagByID)
	r.POST("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.CreateTag)
	r.PUT("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.UpdateTagByID)
	r.DELETE("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.DeleteTag)

	r.Run()
}
