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

func Register(c *gin.Context) {
	var req vm.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handle.RespError(c, http.StatusBadRequest, err)
		return
	}
	rsp, err := authService.Register(c, &req)
	handle.CommonRsp(c, rsp, err)
}

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

	token, err := middleware.GenerateToken(userID, req.Username)
	if err != nil {
		handle.RespError(c, http.StatusInternalServerError, err)
		return
	}
	rsp.Token = token
	handle.CommonRsp(c, rsp, nil)
}

func GetUserDevices(c *gin.Context) {
	userID := c.GetInt64("user_id")
	rsp, err := authService.GetUserDevices(c, userID)
	handle.CommonRsp(c, rsp, err)
}
