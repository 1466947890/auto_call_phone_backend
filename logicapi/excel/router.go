package excel

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdmin(router *gin.RouterGroup) {
	router.POST("/upload_excel", UploadPhoneExcel)
	router.POST("/get_phones", GetPhones)
	router.POST("/get_devices", GetDevices)
	router.POST("/delete_phone", DeletePhone)
	router.POST("/update_phone", UpdatePhone)
}

func RegisterClient(router *gin.RouterGroup) {
	router.POST("/get_client_phone", GetClentPhone)
	router.POST("/update_phone_status", UpdateClientPhoneStatus)
	router.POST("/register_device", RegisterDevice)
}
