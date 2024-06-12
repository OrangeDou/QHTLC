package main

import (
	model "doctorProject/model/server"
	utils "doctorProject/utils/init"

	"github.com/gin-gonic/gin"
)

var (
	logger *utils.Logger
	ctx    *gin.Context
	config = "config.yaml"
)

func main() {
	log := logger.NewLogger()
	route, err := model.RegisterRoute(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	route.Run(":8080")

	err = utils.StartLedger()
	if err != nil {
		log.Error("Fail to start the ledger," + err.Error())
	}
}
