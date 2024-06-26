package http

import (
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/pkg/app"
	"github.com/xian1367/layout-go/pkg/database"
	"github.com/xian1367/layout-go/pkg/gin"
	"github.com/xian1367/layout-go/pkg/redis"
	"github.com/xian1367/layout-go/pkg/shutdown"
	"github.com/xian1367/layout-go/pkg/timer"
	"github.com/xian1367/layout-go/pkg/tracer"
)

// ServerCmd represents the available web sub-command.
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

var Port string

func SetPort(port string) {
	Port = port
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.InitGin()

	//  注册 API 路由
	gin.Routers.RegisterRoutes()

	// Api文档
	if !app.IsProduction() {
		gin.Swagger()
	}

	// 服务连接
	go func() {
		gin.InitServer(Port)
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 cron
		func() {
			timer.Shutdown()
		},

		// 关闭 http server
		func() {
			gin.Shutdown()
		},

		// 关闭 database
		func() {
			database.Shutdown()
		},

		// 关闭 redis
		func() {
			redis.Shutdown()
		},

		// 关闭 tracer
		func() {
			tracer.Shutdown()
		},
	)
}
