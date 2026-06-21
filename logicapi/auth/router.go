package auth

import (
	"auto_call_phone/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterClient(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/devices", middleware.JWTAuth(), GetUserDevices)
}
