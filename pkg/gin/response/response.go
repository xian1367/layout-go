// Package response 响应处理工具
package response

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"reflect"
)

// Failure 错误时返回结构
type Failure struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

type Paging struct {
	Data  interface{} `json:"data"`
	Pager interface{} `json:"pager"`
}

// Data 响应 200 和带 data 键的JSON 数据
// 执行『更新操作』成功后调用，例如更新话题，成功后返回已更新的话题
func Data(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用，例如删除、修改密码、修改手机号
func Success(c *gin.Context) {
	Data(c, gin.H{})
}

// Created 响应 201 和 JSON 数据
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// Abort 响应异常，未传参 msg 时使用默认消息
func Abort(c *gin.Context, code int, msg ...string) {
	c.AbortWithStatusJSON(code, Failure{
		Message: defaultMessage("系统异常", msg...),
	})
}

// Abort400 响应 400，未传参 msg 时使用默认消息
// 在解析用户请求，请求的格式或者方法不符合预期时调用
func Abort400(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Failure{
		Message: defaultMessage("请求错误", msg...),
	})
}

// Abort404 响应 404，未传参 msg 时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, Failure{
		Message: defaultMessage("数据不存在", msg...),
	})
}

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, Failure{
		Message: defaultMessage("无权限", msg...),
	})
}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Failure{
		Message: defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}

// Unauthorized 响应 401，未传参 msg 时使用默认消息
// 登录失败、jwt 解析失败时调用
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Failure{
		Message: defaultMessage("登录失败", msg...),
	})
}

// Error 响应 404 或 422，未传参 msg 时使用默认消息
// 处理请求时出现错误 err，会附带返回 error 信息，如登录错误、找不到 ID 对应的 Model
func Error(c *gin.Context, err error) {
	// error 类型为『数据库未找到内容』
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Failure{
		Message: defaultMessage("请求处理失败", err.Error()),
	})
}

// ValidationError 处理表单验证不通过的错误，返回的 JSON 示例：
//
//	{
//	    Errors: {
//	        "mobile": "手机号长度必须为 11 位的数字"
//	    },
//	    Message: "手机号长度必须为 11 位的数字"
//	}
func ValidationError(c *gin.Context, errors map[string]string) {
	//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Failure{
	//	"message": "表单验证错误",
	//	"errors":  errors,
	//})
	keys := reflect.ValueOf(errors).MapKeys()
	key := keys[rand.Intn(len(keys))].Interface()
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Failure{
		Message: errors[key.(string)],
		Errors:  errors,
	})
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
