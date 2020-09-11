package controllers

import "github.com/gin-gonic/gin"

// Controller definisi interface yang digunakan pada controller
type Controller interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
