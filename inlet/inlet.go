package inlet

import (
	"github.com/spf13/cobra"
	"os"
)

// ConfigPath 存储全局选项 --path 根目录
var ConfigPath string

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(
		&ConfigPath,
		"path",
		"p",
		"./config/setting.yaml",
		"example: --path=./config/setting.yaml",
	)
}

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := ""
	if len(os.Args[1:]) > 0 {
		firstArg = os.Args[1:][0]
	}
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
