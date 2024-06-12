package server

import (
	"doctorProject/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(ctx *gin.Context) (*gin.Engine, error) {
	r := gin.Default()
	//page
	r1 := r.Group("/index")
	{
		r1.GET("/indexPage", handler.HomeIndexHandler)
	}
	return r, nil
}
