package user

import (
	"github.com/xian137/layout-go/database/dao/gen"
	"github.com/xian137/layout-go/database/model"
	"github.com/xian137/layout-go/pkg/database"
	"gorm.io/gorm"
)

const TableName = "users"

func (User) TableName() string {
	return TableName
}

type User struct {
	model.BaseModel
	gen.User
}

func (user *User) Create(tx ...*gorm.DB) {
	db := database.DB
	if len(tx) > 0 {
		db = tx[0]
	}
	db.Create(&user)
}

func (user *User) Save(tx ...*gorm.DB) (rowsAffected int64) {
	db := database.DB
	if len(tx) > 0 {
		db = tx[0]
	}
	result := db.Save(&user)
	return result.RowsAffected
}

func (user *User) Delete(tx ...*gorm.DB) (rowsAffected int64) {
	db := database.DB
	if len(tx) > 0 {
		db = tx[0]
	}
	result := db.Delete(&user)
	return result.RowsAffected
}
