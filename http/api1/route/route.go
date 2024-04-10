package route

import (
	"github.com/xian1367/layout-go/pkg/gin"
	"github.com/xian1367/layout-go/pkg/gin/middleware"
)

type Routes struct{}

// RegisterRoutes 注册 API 相关路由
func (r Routes) RegisterRoutes() {
	api := gin.Engine.Group("api")

	api.Use(middleware.LimitIP("300-H"))
	{
		testRoutes(api)
		userRoutes(api)
	}
}
