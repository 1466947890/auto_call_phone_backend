package main

import (
	"auto_call_phone/logicapi/phone"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库初始化
	fmt.Println("init mysql .....")
}

func initGin() {
	router := gin.Default()
	router.GET("/upload_excel", phone.GetPhoneList)
}
