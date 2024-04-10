package service

import (
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/cmd/service/user"
)

func init() {
	// 注册 make 的子命令
	CmdService.AddCommand(
		user.CmdServiceUser,
	)
}

var CmdService = &cobra.Command{
	Use:   "service",
	Short: "业务层命令",
}
