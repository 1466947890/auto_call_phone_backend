package excel

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdmin(router *gin.RouterGroup) {
	router.POST("/upload_excel", UploadPhoneExcel)
	router.POST("/get_phones", GetPhones)
	router.POST("/delete_phone", DeletePhone)
	router.POST("/update_phone", UpdatePhone)
}
