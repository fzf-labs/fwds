package ratelimit

import (
	"sync"
	"time"
)

//
// @Description:  单机版计数器
//
type counter struct {
	rate  int           //限制值
	count int           //计数值
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	lock  sync.Mutex    //互斥锁
}

// NewCounter 单机版计数器 构造函数
func NewCounter(rate int, cycle time.Duration) *counter {
	return &counter{
		rate:  rate,
		count: 0,
		begin: time.Now(),
		cycle: cycle,
	}
}

// Allow 是否允许通过
func (c *counter) Allow() bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.count > c.rate-1 {
		now := time.Now()
		//已经超过时间周期了
		if now.Sub(c.begin) >= c.cycle {
			c.begin = now
			c.count = 0
			return true
		}
		return false
	} else {
		c.count++
		return true
	}
}
