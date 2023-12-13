package seed

import (
	"github.com/xian137/layout-go/database/seeder"
	"github.com/xian137/layout-go/pkg/console"
	"github.com/xian137/layout-go/pkg/seed"

	"github.com/spf13/cobra"
)

var CmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // 只允许 1 个参数
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeder.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		if len(seed.GetSeeder(name).Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
