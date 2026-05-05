package excel

import (
	"auto_call_phone/logicapi/excel/service"
	"auto_call_phone/logicapi/excel/viewmodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

var excelService *service.ExcelService

func init() {
	excelService = service.NewExcelService()
}

// @Summary Excel上传电话号码
// @Description Excel上传电话号码
// @Tags Excel模块
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "要上传的 Excel 文件"
// @Param user_id formData string false "用户 ID"
// @Success 200 {object} viewmodels.UploadPhoneExcelRsp
// @Router /upload_excel [POST]
func UploadPhoneExcel(c *gin.Context) {

	var req viewmodels.UploadPhoneExcelReq
	// 接收文件
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	file, err := req.File.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	defer file.Close()
	rsp, err := excelService.UploadPhoneExcel(c, nil, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, rsp)

}
