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

// @Summary 获取全部设备列表
// @Description 获取已注册的全部设备唯一标识列表
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.GetDevicesReq true "查询参数"
// @Success 200 {object} viewmodels.GetDevicesRsp
// @Router /v1/admin/excel/get_devices [POST]
func GetDevices(c *gin.Context) {
	var req viewmodels.GetDevicesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.GetDevices(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary 软删除电话号码
// @Description 根据ID软删除电话号码
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.DeletePhoneReq true "删除参数"
// @Success 200 {object} viewmodels.DeletePhoneRsp
// @Router /v1/admin/excel/delete_phone [POST]
func DeletePhone(c *gin.Context) {
	var req viewmodels.DeletePhoneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.DeletePhone(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary 编辑电话号码
// @Description 根据ID编辑电话号码的备注与状态
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.UpdatePhoneReq true "编辑参数"
// @Success 200 {object} viewmodels.UpdatePhoneRsp
// @Router /v1/admin/excel/update_phone [POST]
func UpdatePhone(c *gin.Context) {
	var req viewmodels.UpdatePhoneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.UpdatePhone(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary C端获取待处理电话号码
// @Description 获取指定设备的待处理电话号码数组
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.GetClientPhoneReq true "查询参数"
// @Success 200 {object} viewmodels.GetClentPhoneRsp
// @Router /v1/client/excel/get_client_phone [POST]
func GetClentPhone(c *gin.Context) {
	var req viewmodels.GetClientPhoneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.GetClientPhone(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary C端更新电话号码状态
// @Description 根据电话号码更新当前状态，仅支持 C 端调用
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.UpdateClientPhoneStatusReq true "更新参数"
// @Success 200 {object} viewmodels.UpdateClientPhoneStatusRsp
// @Router /v1/client/excel/update_phone_status [POST]
func UpdateClientPhoneStatus(c *gin.Context) {
	var req viewmodels.UpdateClientPhoneStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.UpdateClientPhoneStatus(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}

// @Summary 绑定设备到用户
// @Description 将指定设备ID绑定到用户，后台管理接口
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.BindDeviceReq true "绑定参数"
// @Success 200 {object} viewmodels.BindDeviceRsp
// @Router /v1/admin/excel/bind_device [POST]
func BindDevice(c *gin.Context) {
	var req viewmodels.BindDeviceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.BindDevice(c, nil, &req)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 概览统计
// @Description 获取系统概览统计，包括设备数、电话数、用户数及各状态分布
// @Tags Excel模块
// @Produce application/json
// @Success 200 {object} viewmodels.OverviewRsp
// @Router /v1/admin/excel/overview [GET]
func GetOverview(c *gin.Context) {
	rsp, err := excelService.GetOverview(c, nil)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 设备统计列表
// @Description 获取所有设备的统计信息，包括绑定用户及各状态电话数量
// @Tags Excel模块
// @Produce application/json
// @Success 200 {object} viewmodels.GetDevicesStatsRsp
// @Router /v1/admin/excel/devices_stats [GET]
func GetDevicesStats(c *gin.Context) {
	rsp, err := excelService.GetDevicesStats(c, nil)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 注册设备唯一标识
// @Description 注册设备唯一标识，如果设备已存在则不重复注册
// @Tags Excel模块
// @Accept application/json
// @Produce application/json
// @Param req body viewmodels.RegisterDeviceReq true "注册参数"
// @Success 200 {object} viewmodels.RegisterDeviceRsp
// @Router /v1/client/excel/register_device [POST]
func RegisterDevice(c *gin.Context) {
	var req viewmodels.RegisterDeviceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := excelService.RegisterDevice(c, nil, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	handle.CommonRsp(c, rsp, err)
}
