package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xian137/layout-go/pkg/console"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Args:  cobra.ExactArgs(2), // 只允许且必须传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		// 格式化模型名称，返回一个 Model 对象
		model := makeModelFromString(args[1])

		dir := fmt.Sprintf("http/%s/request/%s/", args[0], model.PackageName)
		// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			console.Error(err.Error())
		}

		// 基于模板创建文件（做好变量替换）
		createFileFromStub(dir+model.PackageName+"_request.go", "request", model)
	},
}
