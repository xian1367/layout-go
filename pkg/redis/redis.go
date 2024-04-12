// Package redis 工具包
package redis

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/xian1367/layout-go/pkg/logger"
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
	// 开启 tracer instrumentation.
	if err := redisotel.InstrumentTracing(rds.Client); err != nil {
		panic(err)
	}

	// 开启 metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rds.Client); err != nil {
		panic(err)
	}
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

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds Client) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorName("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds Client) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorName("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds Client) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorName("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds Client) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorName("Redis", "Del", err.Error())
		return false
	}
	return true
}

// MultiIncr 事务自增
func (rds Client) MultiIncr(key string, expiration int64) string {
	pipe := rds.Client.TxPipeline()
	incr := pipe.Incr(rds.Context, key)
	pipe.Expire(rds.Context, key, time.Duration(expiration)*time.Second)
	_, err := pipe.Exec(rds.Context)
	if err != nil {
		return ""
	}
	inc := strconv.FormatInt(incr.Val(), 10)
	return inc
}
