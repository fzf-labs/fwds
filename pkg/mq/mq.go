package mq

import (
	"fwds/internal/conf"

	"sync"
	"time"
)

var (
	Mq   IMQ
	Once sync.Once
	Lock sync.Mutex
)

var Consumers = map[*Business]Handle{}

// Handle 消费者业务方法
type Handle func(string) error

//不同类型mq的实现接口
type IMQ interface {
	// Publish 生产普通消息
	Publish(b *Business, msg string) error
	// DeferredPublish 生产延时消息
	DeferredPublish(b *Business, msg string, t time.Duration) error
	// Register 注册一个消费者
	Register(b *Business, handle Handle)
	// Listen 消费者监听
	Listen()
}

func Init() IMQ {
	Once.Do(func() {
		switch conf.Conf.Mq.Use {
		case "RocketMqAli":
			Mq = NewRocketAli()
		case "RocketMq":
			Mq = NewRocket()
		case "Nsq":
			Mq = NewNSQ()
		default:
			panic("New Mq Error")
		}
	})
	return Mq
}

func Register(b *Business, handle Handle) {
	Lock.Lock()
	defer Lock.Unlock()
	Consumers[b] = handle
}

func Listen() {
	if conf.Conf.Mq.Switch {
		Mq.Listen()
	}
}
