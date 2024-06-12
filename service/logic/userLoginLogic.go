package logic

import (
	model "doctorProject/model/user"
	utils "doctorProject/utils/init"

	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// FIXME: fix the logic and error
func UserLoggiLogic(ctx *gin.Context, db *gorm.DB, log *utils.Logger) gin.H {

	msg := gin.H{}
	userIn := model.User{}
	//获取参数ß
	userIn.ID = ctx.PostForm("id")
	userIn.Username = ctx.PostForm("username")
	userIn.Password = ctx.PostForm("password")

	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")

	//验证登录，返回信息
	userOne := model.User{}
	db.Where("userOne_name = ?", userIn.ID).First(&userOne)

	if userOne.Password == userIn.Password && userOne.Username == userIn.Username {
		log.Info("Success logging at " + formatTime)
		msg = gin.H{"message": "Succeed to log"}
	} else {

		msg = gin.H{"message": utils.GetErrMsg(1002)}
		log.Info("用户名或密码不正确 at " + formatTime)
	}

	return msg

}
