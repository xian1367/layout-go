package route

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册 API 相关路由
func RegisterRoutes(route *gin.Engine) {
	var api *gin.RouterGroup
	api = route.Group("api")

	testRoutes(api)
	userRoutes(api)
}
