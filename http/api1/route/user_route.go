package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xian137/layout-go/http/api1/controller"
)

func userRoutes(api *gin.RouterGroup) {
	api.Use()
	{
		ct := new(controller.UserController)
		gp := api.Group("/user")
		{
			gp.GET("", ct.Index)
			gp.GET("/:id", ct.Show)
			gp.POST("", ct.Store)
			gp.PUT("/:id", ct.Update)
			gp.DELETE("/:id", ct.Delete)
		}
	}
}
