package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Start inisialisasi route yang digunakan
func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.Run()
}
