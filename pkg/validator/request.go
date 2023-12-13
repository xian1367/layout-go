package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xian137/layout-go/pkg/app"
	"github.com/xian137/layout-go/pkg/gin/response"
	"strings"
)

// VerifyFunc 验证函数类型
type VerifyFunc func(interface{}) map[string]string

// Validate 控制器里调用示例：
//
//	if ok := request.Validate(c, &requests.UserSaveRequest{}, requests.UserSave); !ok {
//	    return
//	}
func Validate(c *gin.Context, obj interface{}, handlers ...VerifyFunc) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		// 2. 表单验证
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Error(c, err)
		}

		// 3. 翻译
		response.ValidationError(c, RemoveTopStruct(errs.Translate(trans)))
		return false
	}

	// 3. 自定义验证
	if len(handlers) > 0 {
		errs := handlers[0](obj)
		if len(errs) > 0 {
			response.ValidationError(c, errs)
			return false
		}
	}

	return true
}

// GetJson 仅绑定解析
func GetJson(c *gin.Context, obj interface{}) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		//response.BadRequest(c, err, "请求参数错误")
		//return false
	}

	return true
}

// RemoveTopStruct 去除以"."及其左部分内容
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, value := range fields {
		res[app.StrSnake(field[:strings.Index(field, ".")])] = value
	}
	return res
}
