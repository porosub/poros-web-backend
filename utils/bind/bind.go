package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Binding struct{}

type BindingInterface interface {
	BindingJSON(c gin.Context, obj interface{}) error
}

func (b *Binding) BindingJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.JSON); err != nil {
		if err := c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind); err != nil {
			return err
		}
		return err
	}
	return nil
}
