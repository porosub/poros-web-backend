package response

import "github.com/gin-gonic/gin"

func Res(c *gin.Context,
	key, value, status, message string,
	httpStatus int,
	data interface{}) {
	c.Header(key, value)
	c.JSON(httpStatus, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}