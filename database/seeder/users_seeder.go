package seeder

import (
	"fmt"
	"github.com/xian137/layout-go/database/factory"
	"github.com/xian137/layout-go/pkg/console"
	"github.com/xian137/layout-go/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factory.MakeUsers(10)

		result := db.Table("users").Create(&users)

		console.ExitIf(result.Error)

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
