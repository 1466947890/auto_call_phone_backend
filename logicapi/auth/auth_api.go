package auth

import (
	"auto_call_phone/common/handle"
	"auto_call_phone/logicapi/auth/service"
	vm "auto_call_phone/logicapi/auth/viewmodels"
	"auto_call_phone/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService *service.AuthService

func init() {
	authService = service.NewAuthService()
}

// @Summary 用户注册
// @Description 用户注册，创建新账号
// @Tags 认证模块
// @Accept application/json
// @Produce application/json
// @Param req body vm.RegisterReq true "注册参数"
// @Success 200 {object} vm.RegisterRsp
// @Router /v1/client/auth/register [POST]
func Register(c *gin.Context) {
	var req vm.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := authService.Register(c, &req)
	handle.CommonRsp(c, rsp, err)
}

// @Summary 用户登录
// @Description 用户登录，成功返回 JWT Token
// @Tags 认证模块
// @Accept application/json
// @Produce application/json
// @Param req body vm.LoginReq true "登录参数"
// @Success 200 {object} vm.LoginRsp
// @Router /v1/client/auth/login [POST]
func Login(c *gin.Context) {
	var req vm.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	userID, rsp, err := authService.Login(c, &req)
	if err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GenerateToken(userID, req.Username, rsp.Role)
	if err != nil {
		handle.RespError(c, http.StatusInternalServerError, err)
		return
	}
	rsp.Token = token
	handle.CommonRsp(c, rsp, nil)
}

// @Summary 获取用户设备列表
// @Description 获取当前登录用户已注册的设备唯一标识列表
// @Tags 认证模块
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} vm.GetUserDevicesRsp
// @Router /v1/client/auth/devices [GET]
func GetUserDevices(c *gin.Context) {
	userID := c.GetInt64("user_id")
	rsp, err := authService.GetUserDevices(c, userID)
	handle.CommonRsp(c, rsp, err)
}
