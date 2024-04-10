// Package make 命令行的 make 命令
package make

import (
	"embed"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/pkg/app"
	"github.com/xian1367/layout-go/pkg/console"
)

// Model 参数解释
// 单个词，用户命令传参，以 User 模型为例：
//   - user
//   - User
//   - users
//   - Users
//
// 示例：
//
//	{
//	    "TableName": "users",
//	    "StructName": "User",
//	    "StructNamePlural": "Users"
//	    "VariableName": "user",
//	    "VariableNamePlural": "users",
//	    "PackageName": "user"
//	}
//
// -
// 两个词或者以上，用户命令传参，以 TopicComment 模型为例：
//   - topic_comment
//   - topic_comments
//   - TopicComment
//   - TopicComments
//
// 示例：
//
//	{
//	    "TableName": "topic_comments",
//	    "StructName": "TopicComment",
//	    "StructNamePlural": "TopicComments"
//	    "VariableName": "topicComment",
//	    "VariableNamePlural": "topicComments",
//	    "PackageName": "topic_comment"
//	}

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀名的文件

//go:embed stubs
var stubsFS embed.FS

// CmdMake 说明 cobra 命令
var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册 make 的子命令
	CmdMake.AddCommand(
		CmdMakeAll,
		CmdMakeCmd,
		CmdMakeModel,
		CmdMakeAPIController,
		CmdMakeRequest,
		CmdMakeMigration,
		CmdMakeFactory,
		CmdMakeSeeder,
		CmdMakeGen,
		CmdMakeRoute,
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = app.StrSingular(strcase.ToCamel(name))
	model.StructNamePlural = app.StrPlural(model.StructName)
	model.TableName = app.StrSnake(model.StructNamePlural)
	model.VariableName = app.StrLowerCamel(model.StructName)
	model.PackageName = app.StrSnake(model.StructName)
	model.VariableNamePlural = app.StrLowerCamel(model.StructNamePlural)
	return model
}

// createFileFromStub 读取 stub 文件并进行变量替换
// 最后一个选项可选，如若传参，应传 map[string]string 类型，作为附加的变量搜索替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...map[string]string) {
	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		for _, variable := range variables {
			for k, v := range variable {
				replaces[k] = v
			}
		}
	}

	// 目标文件已存在
	if app.FileExists(filePath) {
		console.Warning(filePath + " already exists!")
	} else {
		// 读取 stub 模板文件
		modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
		if err != nil {
			console.Exit(err.Error())
		}
		modelStub := string(modelData)

		// 添加默认的替换变量
		replaces["{{VariableName}}"] = model.VariableName
		replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
		replaces["{{StructName}}"] = model.StructName
		replaces["{{StructNamePlural}}"] = model.StructNamePlural
		replaces["{{PackageName}}"] = model.PackageName
		replaces["{{TableName}}"] = model.TableName

		// 对模板内容做变量替换
		for search, replace := range replaces {
			modelStub = strings.ReplaceAll(modelStub, search, replace)
		}

		// 存储到目标文件中
		err = app.FilePut([]byte(modelStub), filePath)
		if err != nil {
			console.Exit(err.Error())
		}

		// 提示成功
		console.Success(fmt.Sprintf("[%s] created.", filePath))
	}
}
