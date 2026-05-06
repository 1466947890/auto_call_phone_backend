package excel

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdmin(router *gin.RouterGroup) {
	router.POST("/upload_excel", UploadPhoneExcel)
	router.POST("/get_phone_list", ReadPhoneExcel)
}
