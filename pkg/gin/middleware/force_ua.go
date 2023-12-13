// Package middlewares Gin 中间件
package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xian137/layout-go/pkg/gin/response"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.Error(c, errors.New("User-Agent 标头未找到"))
			return
		}

		c.Next()
	}
}
