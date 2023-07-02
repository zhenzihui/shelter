package main

import (
	"github.com/gin-gonic/gin"
	"projectsuika.com/shelter/src/model"
	"projectsuika.com/shelter/src/model/user"
	"projectsuika.com/shelter/src/route"
)

func main() {
	r := gin.Default()
	initDB()
	route.AuthRoutes(r)
	route.InitFileOperateRoutes(r)
	r.Run("0.0.0.0:9008")
}

func initDB() {
	model.ConnectDatabase()
	err := model.DB.AutoMigrate(&user.User{})
	if err != nil {
		return
	}
	//创建默认用户
	user.CreateDefault()
}
