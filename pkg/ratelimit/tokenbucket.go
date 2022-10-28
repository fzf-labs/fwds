package ratelimit

import (
	"time"

	"golang.org/x/time/rate"
)

// TokenBucket 单机版令牌桶
type TokenBucket struct {
	maxBucket int
	sec       time.Duration
	limiter   *rate.Limiter
}

func NewTokenBucket(maxBucket int, sec time.Duration) *TokenBucket {
	return &TokenBucket{
		maxBucket: maxBucket,
		sec:       sec,
		limiter:   rate.NewLimiter(rate.Every(sec), maxBucket),
	}
}

func (t *TokenBucket) Allow() bool {
	return t.limiter.Allow()
}
