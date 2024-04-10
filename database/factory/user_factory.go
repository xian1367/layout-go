package factory

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xian1367/layout-go/database/model/user"
)

func MakeUsers(count int) []user.User {
	var objs []user.User

	for i := 0; i < count; i++ {
		userModel := user.User{}
		gofakeit.Name()
		objs = append(objs, userModel)
	}

	return objs
}
