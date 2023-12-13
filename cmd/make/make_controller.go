package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，example: make controller api1 user",
	Args:  cobra.ExactArgs(2), // 至少传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		model := makeModelFromString(args[1])
		// 组建目标目录
		filePath := fmt.Sprintf("http/%s/controller/%s_controller.go", serviceName, model.PackageName)

		// 基于模板创建文件（做好变量替换）
		createFileFromStub(filePath, "controller", model, map[string]string{
			"{{ServiceName}}": serviceName,
		})

		CmdMakeRoute.Run(cmd, args)
	},
}
