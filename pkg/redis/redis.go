// Package redis 工具包
package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"gohub/pkg/logger"
	"sync"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用db 1
var Redis *RedisClient

// ConnectRedis 连接redis数据库， 设置全局的redis对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient 创建一个新的redis连接
func NewClient(address string, password string, username string, db int) *RedisClient {

	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用redis库里的NewClient初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，并且设置 expiration 过期时间
//func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
//	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
//		//logger.Err
//	}
//}
