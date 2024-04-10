package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/xian1367/layout-go/pkg/database"
	"regexp"
	"strings"
)

func RegisterValidation(v *validator.Validate) {
	_ = v.RegisterValidation("mobile", ruleMobile)
	_ = v.RegisterValidation("exists", ruleExists)
	_ = v.RegisterValidation("not_exists", ruleNotExists)
}

func ruleMobile(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile("^1[345789]{1}\\d{9}$")
	return reg.MatchString(fl.Field().String())
}

func ruleExists(fl validator.FieldLevel) bool {
	rng := strings.Split(strings.TrimPrefix(fl.Param(), "exists:"), " ")

	// 第一个参数，表名称，如 users
	tableName := rng[0]
	// 第二个参数，字段名称，如 email 或者 phone
	dbFiled := rng[1]

	// 用户请求过来的数据
	requestValue := fl.Field().String()

	// 查询数据库
	var exists bool
	database.DB.Raw(
		"SELECT EXISTS (?)",
		database.DB.Table(tableName).Select("1").Where(dbFiled+" = ?", requestValue),
	).First(&exists)

	return exists
}

func ruleNotExists(fl validator.FieldLevel) bool {
	rng := strings.Split(strings.TrimPrefix(fl.Param(), "not_exists:"), " ")

	// 第一个参数，表名称，如 users
	tableName := rng[0]
	// 第二个参数，字段名称，如 email 或者 phone
	dbFiled := rng[1]

	// 第三个参数，排除 ID
	var exceptID string
	if len(rng) > 2 {
		exceptID = rng[2]
	}

	// 用户请求过来的数据
	requestValue := fl.Field().String()

	query := database.DB.Table(tableName).
		Select("1").
		Where(dbFiled+" = ?", requestValue)
	// 拼接 SQL

	// 如果传参第三个参数，加上 SQL Where 过滤
	if len(exceptID) > 0 {
		query.Where("id != ?", exceptID)
	}

	// 查询数据库
	var exists bool
	database.DB.Raw("SELECT EXISTS (?)", query).First(&exists)

	return !exists
}
