package excel

import (
	"auto_call_phone/common/handle"
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
// @Param device_id formData string true "设备 ID"
// @Param user_id formData string false "用户 ID"
// @Success 200 {object} viewmodels.UploadPhoneExcelRsp
// @Router /v1/admin/excel/upload_excel [POST]
func UploadPhoneExcel(c *gin.Context) {
	var req viewmodels.UploadPhoneExcelReq
	if err := c.ShouldBind(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.UploadPhoneExcel(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary 获取电话号码列表
// @Description 获取电话号码列表
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.GetPhonesReq true "查询信息"
// @Success 200 {object} viewmodels.GetPhonesRsp
// @Router /v1/admin/excel/get_phones [POST]
func GetPhones(c *gin.Context) {
	var req viewmodels.GetPhonesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.GetPhones(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}
