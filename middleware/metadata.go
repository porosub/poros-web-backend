package middleware

import (
	"github.com/divisi-developer-poros/poros-web-backend/utils/host"
	"github.com/gin-gonic/gin"
)

// MetadataMiddleware ... Middleware declaration
type MetadataMiddleware struct {
}

// Handler ... store metadata such as host
func (m *MetadataMiddleware) Handler(c *gin.Context) {
	host.Host = c.Request.Host
	host.Scheme = "http"
}
