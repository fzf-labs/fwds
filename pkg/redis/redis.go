package redis

import (
	"context"
	"fwds/internal/conf"

	"fwds/pkg/log"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

const ErrRedisNil = redis.Nil

// Init redis 初始化
func Init(cfg *conf.RedisConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		MinIdleConns: cfg.MinIdleConn,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  cfg.PoolTimeout,
	})
	Client.AddHook(&telescope{})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.SugaredLogger.Errorf("[redis] redis ping err: %+v", err)
		panic(err)
	}
}
