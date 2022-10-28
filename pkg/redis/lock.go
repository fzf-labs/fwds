package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	// LockKey redis lock key
	LockKey = "fwds:redis:lock:%s"
	// DefaultTimeout default expire time
	DefaultTimeout = 10 * time.Second
)

// Lock 定义lock结构体
type Lock struct {
	key         string
	redisClient *redis.Client
	token       string
	timeout     time.Duration
	// todo: support retry
	// maxRetries int
}

type Option func(*Lock)

// WithTimeout  过期时间设置
func WithTimeout(expiration time.Duration) Option {
	return func(l *Lock) {
		l.timeout = expiration
	}
}

func NewDefaultLock(key string) *Lock {
	return NewLock(Client, key)
}

// NewLock new a lock instance
func NewLock(conn *redis.Client, key string, opts ...Option) *Lock {
	lock := Lock{
		key:         key,
		redisClient: conn,
		token:       genToken(),
		timeout:     DefaultTimeout,
	}

	for _, o := range opts {
		o(&lock)
	}
	return &lock
}

// Lock 加锁
func (l *Lock) Lock(ctx context.Context) (bool, error) {
	ok, err := l.redisClient.SetNX(ctx, l.GetKey(), l.token, l.timeout).Result()
	if err == redis.Nil {
		err = nil
	}
	return ok, err
}

// Unlock 解锁
// token 一致才会执行删除，避免误删，这里用了lua脚本进行事务处理
func (l *Lock) Unlock(ctx context.Context) error {
	script := "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	_, err := l.redisClient.Eval(ctx, script, []string{l.GetKey()}, l.token).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetKey 获取key
func (l *Lock) GetKey() string {
	return fmt.Sprintf(LockKey, l.key)
}

// genToken 生成token
func genToken() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
