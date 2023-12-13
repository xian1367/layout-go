package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xian137/layout-go/database/model/scope"
	"github.com/xian137/layout-go/pkg/database"
)

func Get(id interface{}) (user User) {
	database.DB.Where("id = ?", id).First(&user)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context) (users []User, paging scope.Paging) {
	paging = scope.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		User{},
	)
	return
}
