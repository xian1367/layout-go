package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeRoute = &cobra.Command{
	Use:   "route",
	Short: "Crate route file, example: make route api1 user",
	Args:  cobra.ExactArgs(2), // 只允许且必须传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]

		model := makeModelFromString(args[1])

		// 组建目标目录
		filePath := fmt.Sprintf("http/%s/route/%s_route.go", serviceName, model.PackageName)

		// 基于模板创建文件（做好变量替换）
		createFileFromStub(filePath, "route", model, map[string]string{
			"{{ServiceName}}": serviceName,
		})
	},
}
