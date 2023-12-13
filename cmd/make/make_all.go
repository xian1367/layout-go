package make

import (
	"github.com/spf13/cobra"
)

var CmdMakeAll = &cobra.Command{
	Use:   "all",
	Short: "Crate all file, example: make all api1 user",
	Args:  cobra.ExactArgs(2), // 至少传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		CmdMakeCmd.Run(cmd, args)
		CmdMakeAPIController.Run(cmd, args)
		CmdMakeFactory.Run(cmd, args)
		CmdMakeGen.Run(cmd, args)
		CmdMakeMigration.Run(cmd, args)
		CmdMakeModel.Run(cmd, args)
		CmdMakeRequest.Run(cmd, args)
		CmdMakeSeeder.Run(cmd, args)
		CmdMakeRoute.Run(cmd, args)
	},
}
