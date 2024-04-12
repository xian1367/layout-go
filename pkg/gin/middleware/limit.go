package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/xian1367/layout-go/pkg/app"
	"github.com/xian1367/layout-go/pkg/gin/limit"
	"github.com/xian1367/layout-go/pkg/gin/response"
)

// LimitIP 全局限流中间件，针对 IP 进行限流
// rate  为格式化字符串，如 "5-S" ，示例:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
func LimitIP(param string) gin.HandlerFunc {
	if app.IsDebug() {
		param = "1000000-H"
	}

	return func(c *gin.Context) {
		// 针对 IP 限流
		key := limit.GetKeyIP(c)
		if ok := limitHandler(c, key, param); !ok {
			return
		}
		c.Next()
	}
}

// LimitPerRoute 限流中间件，用在单独的路由中
func LimitPerRoute(param string) gin.HandlerFunc {
	if app.IsDebug() {
		param = "1000000-H"
	}
	return func(c *gin.Context) {

		// 针对单个路由，增加访问次数
		c.Set("limiter-once", false)

		// 针对 IP + 路由进行限流
		key := limit.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, param); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, param string) bool {

	// 获取超额的情况
	rate, err := limit.CheckRate(c, key, param)
	if err != nil {
		response.Abort500(c)
		return false
	}

	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// 超额
	if rate.Reached {
		// 提示用户超额了
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}

	return true
}
