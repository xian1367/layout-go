package migration

import (
	"github.com/xian1367/layout-go/database/model"
	"github.com/xian1367/layout-go/pkg/migrate"
)

func init() {
	type User struct {
		model.IDField

		Name string `gorm:"type:varchar(255);not null;index"`

		model.CommonTimestampsField
	}

	migrate.Add(&User{})
}
