package main

import (
	"auto_call_phone/data/repo"
	"auto_call_phone/logicapi/auth"
	"auto_call_phone/logicapi/excel"
	"auto_call_phone/logicapi/index"
	"auto_call_phone/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库初始化
	fmt.Println("init mysql .....")
	if err := repo.InitGrom(); err != nil {
		log.Fatalln("fail to init mysql : ", err.Error())
	}
	initGin()
}

func initGin() {
	router := gin.Default()

	// 使用CORS中间件
	router.Use(middleware.CORSMiddleware())

	apiV1Admin := router.Group("/v1/admin")
	excel.RegisterAdmin(apiV1Admin.Group("/excel"))
	apiV1Client := router.Group("/v1/client")
	excel.RegisterClient(apiV1Client.Group("/excel"))
	auth.RegisterClient(apiV1Client.Group("/auth"))
	index.RegisterClient(apiV1Client.Group("/"))
	router.Run(":8000")
}
