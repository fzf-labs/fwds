package mq

import (
	"fwds/internal/mq/test"
	"fwds/pkg/mq"
)

func Init() {
	mq.Init()
	Instance()
	mq.Listen()
}

func Instance() {
	test.Consumer()
}
