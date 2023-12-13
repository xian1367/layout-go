package migrate

import (
	"github.com/spf13/cobra"
	"github.com/xian137/layout-go/database/migration"
	"github.com/xian137/layout-go/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	// 所有 migrate 下的子命令都会执行以下代码
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateRefresh,
		CmdMigrateReset,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migration 下的所有迁移文件
	migration.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unMigrated migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator().Up()
	},
}

var CmdMigrateRollback = &cobra.Command{
	Use: "down",
	// 设置别名 migrate down == migrate rollback
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run: func(cmd *cobra.Command, args []string) {
		migrator().Rollback()
	},
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator().Reset()
	},
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator().Refresh()
	},
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator().Fresh()
	},
}
