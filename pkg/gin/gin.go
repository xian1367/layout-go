package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xian137/layout-go/docs"
	"github.com/xian137/layout-go/pkg/console"
	"github.com/xian137/layout-go/pkg/gin/middleware"
	"github.com/xian137/layout-go/pkg/logger"
	"net/http"
	"time"
)

var Engine *gin.Engine

var Server *http.Server

func InitGin() {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)
	// gin 实例
	Engine = gin.Default()
	// 初始化路由
	setupRoute()
}

func InitServer(port string) {
	// 服务连接
	Server = &http.Server{
		Addr:    ":" + port,
		Handler: Engine,
	}
	if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("CMD server", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}

// setupRoute 路由初始化
func setupRoute() {
	// 注册全局中间件
	registerGlobalMiddleWare()
	//  配置 404 路由
	setup404Handler()
}

func registerGlobalMiddleWare() {
	Engine.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.ForceUA(),
		middleware.Response(),
		middleware.Cors(),
	)
}

func setup404Handler() {
	// 处理 404 请求
	Engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "迷路了?")
	})
}

func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := Server.Shutdown(ctx)
	logger.ErrorIf(err)
}

func Swagger() {
	docs.SwaggerInfo.BasePath = "/v1"
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
