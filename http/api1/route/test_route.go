package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xian1367/layout-go/http/api1/controller"
)

func testRoutes(api *gin.RouterGroup) {
	api.Use()
	{
		Test := new(controller.TestController)
		TestGroup := api.Group("/test")
		{
			TestGroup.Any("", Test.Index)
		}
	}
}
