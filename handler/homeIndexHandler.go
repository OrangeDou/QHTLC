package handler

import (
	"doctorProject/service"
	"doctorProject/service/response"

	"github.com/gin-gonic/gin"
)

func HomeIndexHandler(ctx *gin.Context) {

	response.Success(ctx, service.Index(ctx), "Succeed to log")

}
