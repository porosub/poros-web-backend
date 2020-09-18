package routes

import (
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	"net/http"

	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"
	"github.com/gin-gonic/gin"
)

var (
	TestingHandlers test.Cobs
	TokenMiddleware middleware.TokenMiddleware
)

// Start inisialisasi route yang digunakan
func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})


	r.GET("/guest", TestingHandlers.Guest)
	r.POST("/login", TestingHandlers.Login)
	r.GET("/home", TokenMiddleware.AuthorizeToken, TestingHandlers.Home)

	r.Run()
}
