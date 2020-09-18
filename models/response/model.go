package response

import "github.com/gin-gonic/gin"

type Response struct {
	status  string      `json:"status"`
	message string      `json:"message"`
	data    interface{} `json:"data"`
}

type ResInterface interface {
	CustomResponse(c *gin.Context,
		key, value, status, message string,
		httpStatus int,
		data interface{})
}