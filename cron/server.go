package cron

import (
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/pkg/shutdown"
	"github.com/xian1367/layout-go/pkg/timer"
)

// CmdCron represents the available web sub-command.
var CmdCron = &cobra.Command{
	Use:   "cron",
	Short: "Start cron",
	Run:   runCron,
	Args:  cobra.NoArgs,
}

func runCron(cmd *cobra.Command, args []string) {
	Kernel()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 cron
		func() {
			timer.Shutdown()
		},
	)
}
