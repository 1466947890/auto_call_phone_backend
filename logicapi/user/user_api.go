package user

import (
	"auto_call_phone/common/handle"
	"auto_call_phone/logicapi/user/service"
	vm "auto_call_phone/logicapi/user/viewmodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService *service.UserService

func init() {
	userService = service.NewUserService()
}

// @Summary 用户列表
// @Description 分页获取用户列表，支持用户名搜索，返回每人绑定的设备数和电话号码总数
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param req body vm.GetUsersListReq true "查询参数"
// @Success 200 {object} vm.GetUsersListRsp
// @Router /v1/admin/users [POST]
func GetUsersList(c *gin.Context) {
	var req vm.GetUsersListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := userService.GetUsersList(c, nil, &req)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 设备列表
// @Description 分页获取设备列表，支持按绑定状态（all/bound/unbound）筛选
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param req body vm.GetDevicesListReq true "查询参数"
// @Success 200 {object} vm.GetDevicesListRsp
// @Router /v1/admin/devices [POST]
func GetDevicesList(c *gin.Context) {
	var req vm.GetDevicesListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := userService.GetDevicesList(c, nil, &req)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 解绑设备
// @Description 将设备的 user_id 置为 NULL，解除设备与用户的绑定关系
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param req body vm.UnbindDeviceReq true "解绑参数"
// @Success 200 {object} vm.UnbindDeviceRsp
// @Router /v1/admin/devices/unbind [POST]
func UnbindDevice(c *gin.Context) {
	var req vm.UnbindDeviceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := userService.UnbindDevice(c, nil, &req)
	handle.CommonRsp(c, rsp, err)
}
