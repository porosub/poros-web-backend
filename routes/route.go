package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/divisi-developer-poros/poros-web-backend/controllers/user"
	"net/http"
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

	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.Get)

	r.POST("/users", userController.Create)

	r.PUT("/users/:id", userController.Update)

	r.DELETE("/users/:id", userController.Delete)

	r.Run()
}
