package test

import (
	"encoding/json"
	"fmt"
	"fwds/pkg/log"
	"fwds/pkg/mq"
)

func Producer(msg string) {
	prefix := "【消息队列-测试】"
	body, _ := json.Marshal(msg)
	business, err := mq.GetBusiness("test")
	if err != nil {
		log.SugaredLogger.Errorf(prefix+"MQ业务key获取失败,err:%v", err)
		return
	}
	err = mq.Mq.Publish(business, string(body))
	if err != nil {
		log.SugaredLogger.Errorf(prefix+"MQ消息生产失败,err:%v", err)
		return
	}
}

func Consumer() {
	prefix := "【消息队列-测试】"
	business, err := mq.GetBusiness("test")
	if err != nil {
		log.SugaredLogger.Errorf(prefix+"MQ业务key获取失败,err:%v", err)
		return
	}
	mq.Register(business, do)
	log.SugaredLogger.Info(prefix + "消费者,注册成功")
}

func do(string2 string) error {
	fmt.Println(string2)

	return nil
}
