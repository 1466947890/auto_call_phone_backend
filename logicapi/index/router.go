package index

import (
	"auto_call_phone/common/handle"

	"github.com/gin-gonic/gin"
)

func RegisterClient(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		handle.Resp(ctx, 200, "ok", nil)
	})
}
