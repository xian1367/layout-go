// Package middlewares Gin 中间件
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xian1367/layout-go/pkg/gin/response"
	"github.com/xian1367/layout-go/pkg/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("Token认证失败"))
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", claims.UserID)

		c.Next()
	}
}
