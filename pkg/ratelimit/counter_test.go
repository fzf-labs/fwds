package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func Test_counter_Allow(t *testing.T) {
	counter := NewCounter(10, time.Minute)
	i := 1
	for {
		allow := counter.Allow()
		fmt.Println(i, allow)
		i++
		if !allow {
			return
		}
	}
}
