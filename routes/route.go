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
	r.GET("/user/:id", userController.Get)

	r.POST("/user", userController.Create)

	r.PUT("/user/:id", userController.Update)

	r.DELETE("/user/:id", userController.Delete)

	r.Run()
}
