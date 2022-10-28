package ratelimit

import (
	"sync"
	"time"
)

//
// @Description: 单机版滑动窗口
//
type slidingWindow struct {
	max          int           //限制值
	slotDuration time.Duration //插槽的时间长度
	winDuration  time.Duration //整个窗口的时间长度
	slotNum      int           //插槽的个数
	window       []*slot
	mu           sync.Mutex //锁
}

func NewSlidingWindow(max int, slotDuration time.Duration, winDuration time.Duration) *slidingWindow {
	return &slidingWindow{
		max:          max,
		slotDuration: slotDuration,
		winDuration:  winDuration,
		slotNum:      int(winDuration / slotDuration),
	}
}

//
// @Description: 时间插槽
//
type slot struct {
	begin time.Time //这个插槽的起始时间
	count int       //计数
}

// Allow 是否允许
func (s *slidingWindow) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	//1.将过期的插槽移除
	now := time.Now()
	timeoutOffset := -1
	for i, ws := range s.window {
		if ws.begin.Add(s.winDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		s.window = s.window[timeoutOffset+1:]
	}
	//2.判断请求
	var result bool
	if s.count() < s.max {
		result = true
	}
	//3.计数
	var lastSlot *slot
	if len(s.window) > 0 {
		lastSlot = s.window[len(s.window)-1]
		if lastSlot.begin.Add(s.slotDuration).Before(now) {
			// 如果当前时间已经超过这个时间插槽的跨度，那么新建一个时间插槽
			s.addSlot(now)
		} else {
			lastSlot.count++
		}
	} else {
		//不存在插槽时
		s.addSlot(now)
	}

	return result
}

// addSlot 添加插槽
func (s *slidingWindow) addSlot(now time.Time) {
	lastSlot := &slot{
		begin: now,
		count: 1,
	}
	s.window = append(s.window, lastSlot)
}

// count 统计总数
func (s *slidingWindow) count() int {
	var count int
	for _, ws := range s.window {
		count += ws.count
	}
	return count
}
