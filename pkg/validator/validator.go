// Package validator 处理请求数据和表单验证
package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// 定义一个全局翻译器T
var trans ut.Translator

// init 初始化
func init() {
	locale := "zh"
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册自定义验证器
		RegisterValidation(v)
		//注册一个获取json的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("trans")
			if name == "" {
				name = strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				name = field.Name
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, _ = uni.GetTranslator(locale)
		if ok {
			// 注册翻译器
			_ = zhTranslations.RegisterDefaultTranslations(v, trans)
		}

		//注册自定义翻译
		RegisterTranslation(v)
	}
}
