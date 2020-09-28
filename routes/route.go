package routes

import (
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	userController "github.com/divisi-developer-poros/poros-web-backend/controllers/user"
	"net/http"
	"github.com/gin-gonic/gin"
	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"


)

var (
	TestingHandlers test.Cobs
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

	r.Run()
}
