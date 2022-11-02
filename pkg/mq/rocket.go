package mq

import (
	"context"
	"fwds/internal/conf"
	"fwds/pkg/log"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

// DelayTime 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
var DelayTime = map[int]int{
	1:  1,
	2:  5,
	3:  10,
	4:  30,
	5:  60,
	6:  120,
	7:  180,
	8:  240,
	9:  300,
	10: 360,
	11: 420,
	12: 480,
	13: 540,
	14: 600,
	15: 1200,
	16: 1800,
	17: 3600,
	18: 7200,
}

func NewRocket() *Rocket {
	return &Rocket{}
}

type Rocket struct {
}

func (r *Rocket) Publish(b *Business, msg string) error {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.Conf.Mq.RocketMq.Endpoint})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		log.SugaredLogger.Errorf("rocket mq 生产者启动失败")
		panic(err)
	}
	message := &primitive.Message{
		Topic: b.Topic,
		Body:  []byte(msg),
	}
	message.WithTag(b.Tag)
	message.WithKeys([]string{utils.GetUUID()})
	ret, err := p.SendSync(context.Background(), message)
	if err != nil {
		log.SugaredLogger.Errorf("send message error: %s", err)
		return GeneralMessageDeliveryFailed
	}
	log.SugaredLogger.Infof("rocketmq 普通消息生产成功,topic:%s,tag:%s,MsgID:%s,Body:%s", b.Topic, b.Tag, ret.MsgID, msg)
	err = p.Shutdown()
	if err != nil {
		log.SugaredLogger.Errorf("shutdown producer error: %s", err)
	}
	return nil
}

func (r *Rocket) DeferredPublish(b *Business, msg string, t time.Duration) error {
	level := r.GetDelayTimeLevel(t)
	if level == 0 {
		return DelayLevelError
	}
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.Conf.Mq.RocketMq.Endpoint})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		log.SugaredLogger.Errorf("rocket mq 生产者启动失败")
		panic(err)
	}
	message := &primitive.Message{
		Topic: b.Topic,
		Body:  []byte(msg),
	}
	message.WithTag(b.Tag)
	message.WithDelayTimeLevel(level)
	ret, err := p.SendSync(context.Background(), message)
	if err != nil {
		log.SugaredLogger.Errorf("send message error: %s", err)
		return GeneralMessageDeliveryFailed
	}
	log.SugaredLogger.Infof("rocketmq 延时消息生产成功,topic:%s,tag:%s,MsgID:%s,Body:%s", b.Topic, b.Tag, ret.MsgID, msg)
	err = p.Shutdown()
	if err != nil {
		log.SugaredLogger.Errorf("shutdown producer error: %s", err)
	}
	return nil
}

func (r *Rocket) Register(b *Business, handle Handle) {
	Lock.Lock()
	defer Lock.Unlock()
	Consumers[b] = handle
}

func (r *Rocket) Listen() {
	if len(Consumers) > 0 {
		for business, handle := range Consumers {
			go r.do(business, handle)
		}
	}
	log.SugaredLogger.Infof("rocketmq消费者监听成功,共%d个消费者", len(Consumers))
}

func (r *Rocket) GetDelayTimeLevel(t time.Duration) int {
	second := int(t.Seconds())
	for k, v := range DelayTime {
		if v == second {
			return k
		}
	}
	return 0
}
func (r *Rocket) do(b *Business, handle Handle) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(b.GroupId),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.Conf.Mq.RocketMq.Endpoint})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset), // 选择消费时间(首次/当前/根据时间)
		consumer.WithConsumerModel(consumer.Clustering),                // 消费模式(集群消费:消费完其他人不能再读取/广播消费：所有人都能读)
	)
	err := c.Subscribe(
		b.Topic,
		consumer.MessageSelector{
			Type:       consumer.TAG,
			Expression: b.Tag,
		},

		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				err := handle(string(msgs[i].Body))
				if err != nil {
					log.SugaredLogger.Errorf("rocketmq 消息业务处理失败 topic:%v,tag:%v,group_id:%v,body:%v, err:%v", b.Topic, b.Tag, b.GroupId, string(msgs[i].Body), err)
					return 0, err
				}
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		log.SugaredLogger.Errorf("rocket mq Subscribe err:%v", err)
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		log.SugaredLogger.Errorf("rocket mq Start err:%v", err)
		panic(err)
	}
	defer func(c rocketmq.PushConsumer) {
		_ = c.Shutdown()
	}(c)
	select {}
}
