package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xian137/layout-go/bootstrap"
	"github.com/xian137/layout-go/cmd/make"
	"github.com/xian137/layout-go/cmd/migrate"
	"github.com/xian137/layout-go/cmd/seed"
	"github.com/xian137/layout-go/cmd/service"
	"github.com/xian137/layout-go/config"
	"github.com/xian137/layout-go/inlet"
	"github.com/xian137/layout-go/pkg/console"
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
		make.CmdMake,
		migrate.CmdMigrate,
		seed.CmdSeed,
		service.CmdService,
	)

	// 注册全局参数，--config
	inlet.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
