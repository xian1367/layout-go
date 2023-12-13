package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterTranslation(v *validator.Validate) {
	_ = v.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0} 不是一个手机号", true)
	}, func(ut ut.Translator, fl validator.FieldError) string {
		t, _ := ut.T("mobile", fl.Field())
		return t
	})

	_ = v.RegisterTranslation("exists", trans, func(ut ut.Translator) error {
		return ut.Add("exists", "{0} {1} 不存在", true)
	}, func(ut ut.Translator, fl validator.FieldError) string {
		t, _ := ut.T("exists", fl.Field(), fl.Value().(string))
		return t
	})

	_ = v.RegisterTranslation("not_exists", trans, func(ut ut.Translator) error {
		return ut.Add("not_exists", "{0} {1} 已存在", true)
	}, func(ut ut.Translator, fl validator.FieldError) string {
		t, _ := ut.T("not_exists", fl.Field(), fl.Value().(string))
		return t
	})
}
