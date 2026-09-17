package middleware

import (
	"auto_call_phone/common/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware 返回配置好的CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	corsCfg := config.App.CORS
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			for _, allowOrigin := range corsCfg.AllowOrigins {
				if origin == allowOrigin {
					return true
				}
			}
			for _, prefix := range corsCfg.AllowOriginPrefixes {
				if strings.HasPrefix(origin, prefix) {
					return true
				}
			}
			for _, contains := range corsCfg.AllowOriginContains {
				if strings.Contains(origin, contains) {
					return true
				}
			}
			return false
		},
		AllowMethods:     corsCfg.AllowMethods,
		AllowHeaders:     corsCfg.AllowHeaders,
		ExposeHeaders:    corsCfg.ExposeHeaders,
		AllowCredentials: corsCfg.AllowCredentials,
	})
}
