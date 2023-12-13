// Package redis 工具包
package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xian137/layout-go/pkg/logger"
	"strconv"
	"sync"
	"time"
)

// Client Redis 服务
type Client struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var Redis *Client

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
		// 测试一下连接
		err := Ping()
		logger.ErrorIf(err)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int) *Client {
	// 初始化自定的 RedisClient 实例
	rds := &Client{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	return rds
}

// Ping 用以测试 redis 连接是否正常
func Ping() error {
	_, err := Redis.Client.Ping(Redis.Context).Result()
	return err
}

func Shutdown() {
	err := Redis.Client.Close()
	logger.ErrorIf(err)
}

// MultiIncr 事务自增
func MultiIncr(key string, expiration int64) string {
	pipe := Redis.Client.TxPipeline()
	incr := pipe.Incr(Redis.Context, key)
	pipe.Expire(Redis.Context, key, time.Duration(expiration)*time.Second)
	_, err := pipe.Exec(Redis.Context)
	if err != nil {
		return ""
	}
	inc := strconv.FormatInt(incr.Val(), 10)
	return inc
}
