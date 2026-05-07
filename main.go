package main

import (
	"auto_call_phone/data/repo"
	"auto_call_phone/logicapi/excel"
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
	apiV1Admin := router.Group("/v1/admin")
	excel.RegisterAdmin(apiV1Admin.Group("/excel"))
	router.Run(":80")
}
