package main

import (
	"auto_call_phone/data/repo"
	_ "auto_call_phone/docs"
	"auto_call_phone/logicapi/auth"
	"auto_call_phone/logicapi/excel"
	"auto_call_phone/logicapi/index"
	"auto_call_phone/logicapi/user"
	"auto_call_phone/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 自动打电话后端 API
// @version 1.0
// @description 自动打电话系统后端接口文档
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	apiV1Admin.Use(middleware.JWTAuth(), middleware.AdminAuth())
	excel.RegisterAdmin(apiV1Admin.Group("/excel"))
	user.RegisterAdmin(apiV1Admin.Group("/"))
	apiV1Client := router.Group("/v1/client")
	excel.RegisterClient(apiV1Client.Group("/excel"))
	auth.RegisterClient(apiV1Client.Group("/auth"))
	index.RegisterClient(apiV1Client.Group("/"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8000")
}
