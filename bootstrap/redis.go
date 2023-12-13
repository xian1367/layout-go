package bootstrap

import (
	"fmt"
	"github.com/xian137/layout-go/config"
	"github.com/xian137/layout-go/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.Get().Redis.Host, config.Get().Redis.Port),
		config.Get().Redis.Username,
		config.Get().Redis.Password,
		config.Get().Redis.DBName,
	)
}
