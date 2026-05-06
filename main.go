package main

import (
	"auto_call_phone/logicapi/excel"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库初始化
	fmt.Println("init mysql .....")
	initGin()
}

func initGin() {
	router := gin.Default()

	adminRouter := router.Group("/v1/admin")
	excel.RegisterAdmin(adminRouter)

	router.Run(":80")
}
