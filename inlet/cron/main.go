package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/bootstrap"
	"github.com/xian1367/layout-go/config"
	"github.com/xian1367/layout-go/cron"
	"github.com/xian1367/layout-go/inlet"
	"github.com/xian1367/layout-go/pkg/console"
	"os"
)

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get().App.Name,
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --config 参数
			config.InitConfig(inlet.ConfigPath)
			// 初始化时区
			config.InitTime()

			// 初始化 定时器
			bootstrap.SetupTimer()

			// 初始化 日志
			bootstrap.SetupLogger("cron")

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// UI-Log
			bootstrap.UI()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cron.CmdCron,
	)

	// 配置默认运行服务
	inlet.RegisterDefaultCmd(rootCmd, cron.CmdCron)

	// 注册全局参数，--config
	inlet.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
