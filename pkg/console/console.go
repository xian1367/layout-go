// Package console 命令行辅助方法
package console

import (
	"github.com/fatih/color"
	"os"
)

// Success 打印一条成功消息，绿色输出
func Success(msg string) {
	color.Green(msg)
}

// Error 打印一条报错消息，红色输出
func Error(msg string) {
	color.Red(msg)
}

// Warning 打印一条提示消息，黄色输出
func Warning(msg string) {
	color.Yellow(msg)
}

// Exit 打印一条报错消息，并退出 os.Exit(1)
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}
