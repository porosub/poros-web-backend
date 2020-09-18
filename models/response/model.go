package response

import "github.com/gin-gonic/gin"

type Response struct {}

type ResInterface interface {
	CustomResponse(c *gin.Context,
		key, value, status, message string,
		httpStatus int,
		data interface{})
}