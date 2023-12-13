package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xian137/layout-go/pkg/console"
	"os"
)

var CmdMakeCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd buckup_database",
	Args:  cobra.ExactArgs(2), // 只允许且必须传 1 个参数
	Run: func(cmd *cobra.Command, args []string) {
		// 格式化模型名称，返回一个 Model 对象
		model := makeModelFromString(args[0])

		dir := fmt.Sprintf("cmd/service/%s/", args[0])
		// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			console.Error(err.Error())
		}

		// 从模板中创建文件（做好变量替换）
		createFileFromStub(dir+fmt.Sprintf("%s.go", args[1]), "cmd", model)

		// 友好提示
		console.Success("command name:" + model.PackageName)
		console.Success("command variable name: cmd.Cmd" + model.StructName)
		console.Warning("please edit main.go's app.Commands slice to register command")
	},
}
