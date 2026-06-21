package index

import (
	"auto_call_phone/common/handle"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查
// @Description 服务健康检查，返回 ok 表示服务正常运行
// @Tags 系统模块
// @Produce application/json
// @Success 200 {object} handle.Response
// @Router /v1/client/ [GET]
func RegisterClient(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		handle.Resp(ctx, 200, "ok", nil)
	})
}
