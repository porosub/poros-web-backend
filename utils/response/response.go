package response

import "github.com/gin-gonic/gin"

// Response ... Response struct declaration
type Response struct{}

// ResInterface ... Res interface declaration
type ResInterface interface {
	CustomResponse(c *gin.Context,
		key, value, status, message string,
		httpStatus int,
		data interface{})
}

// CustomResponse ... Implement custom response
func (r *Response) CustomResponse(c *gin.Context,
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
