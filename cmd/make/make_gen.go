package make

import (
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go/pkg/database"
	"gorm.io/gen"
)

var CmdMakeGen = &cobra.Command{
	Use:   "gen",
	Short: "Generate file and code, example: make gen",
	Run: func(cmd *cobra.Command, args []string) {
		g := gen.NewGenerator(gen.Config{
			OutPath:      "database/dao/query",
			ModelPkgPath: "model_gen",
		})

		g.UseDB(database.DB)

		fieldOpts := []gen.ModelOpt{
			gen.FieldIgnore("id", "created_at", "updated_at", "deleted_at"),
		}

		allModel := g.GenerateAllTable(fieldOpts...)

		g.ApplyBasic(allModel...)

		g.Execute()
	},
}
