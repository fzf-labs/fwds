package ratelimit

import (
	rl "go.uber.org/ratelimit"
)

//漏桶算法
type leakyBucket struct {
	rate    int
	limiter rl.Limiter
}

func NewLeakyBucket(rate int) *leakyBucket {
	return &leakyBucket{
		rate:    rate,
		limiter: rl.New(rate),
	}
}

//func (t *leakyBucket) Allow() bool {
//	return t.limiter.Take()
//}
