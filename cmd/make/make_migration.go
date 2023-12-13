package make

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xian137/layout-go/pkg/console"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration user",
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	Run: func(cmd *cobra.Command, args []string) {
		model := makeModelFromString(args[0])

		filePath := fmt.Sprintf("database/migration/%s_migration.go", model.PackageName)

		createFileFromStub(filePath, "migration", model)

		console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
	},
}
