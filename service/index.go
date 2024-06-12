package service

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) gin.H {
	return gin.H{
		"message": "Hello World",
	}
}
