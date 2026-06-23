package user

import "github.com/gin-gonic/gin"

func RegisterAdmin(router *gin.RouterGroup) {
	router.POST("/users", GetUsersList)
	router.POST("/devices", GetDevicesList)
	router.POST("/devices/unbind", UnbindDevice)
}
