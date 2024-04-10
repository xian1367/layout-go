// Package limit 处理限流逻辑
package limit

import (
	"github.com/xian1367/layout-go/config"
	"github.com/xian1367/layout-go/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
	"github.com/xian1367/layout-go/pkg/redis"
)

// GetKeyIP 获取 Limiter 的 Key，IP
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP Limitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

// CheckRate 检测请求是否超额
func CheckRate(c *gin.Context, key string, formatted string) (limiter.Context, error) {
	// 实例化依赖的 limiter 包的 limiter.Rate 对象
	var context limiter.Context
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		logger.ErrorIf(err)
		return context, err
	}

	// 初始化存储，使用我们程序里共用的 redis.Redis 对象
	store, err := limiterRedis.NewStoreWithOptions(redis.Redis.Client, limiter.StoreOptions{
		// 为 limiter 设置前缀，保持 redis 里数据的整洁
		Prefix: config.Get().App.Name + ":limiter",
	})
	if err != nil {
		logger.ErrorIf(err)
		return context, err
	}

	// 使用上面的初始化的 limiter.Rate 对象和存储对象
	limiterObj := limiter.New(store, rate)

	// 获取限流的结果
	if c.GetBool("limiter-once") {
		// Peek() 取结果，不增加访问次数
		return limiterObj.Peek(c, key)
	} else {

		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数。
		c.Set("limiter-once", true)

		// Get() 取结果且增加访问次数
		return limiterObj.Get(c, key)
	}
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
