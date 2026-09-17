package handle

import (
	"auto_call_phone/common/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS 配置CORS中间件
func SetupCORS() gin.HandlerFunc {
	corsCfg := config.App.CORS
	return cors.New(cors.Config{
		AllowOrigins:     corsCfg.AllowOrigins,
		AllowMethods:     corsCfg.AllowMethods,
		AllowHeaders:     corsCfg.AllowHeaders,
		ExposeHeaders:    corsCfg.ExposeHeaders,
		AllowCredentials: corsCfg.AllowCredentials,
	})
}
