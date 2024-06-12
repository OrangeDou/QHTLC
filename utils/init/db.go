package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const configPath = "config.yaml"

func ConnectDB() *gorm.DB {
	log := Logger{}
	l := log.NewLogger()
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		l.Error("Failed to connect the config file, err=" + err.Error())
	}

	user := viper.GetString("mysqlRemote.user")
	password := viper.GetString("mysqlRemote.password")
	host := viper.GetString("mysqlRemote.host")
	port := viper.GetString("mysqlRemote.port")
	database := viper.GetString("mysqlRemote.database")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	fmt.Println(user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		l.Error("Failed to connect the mysql, err=" + err.Error())

	}
	return db
}
