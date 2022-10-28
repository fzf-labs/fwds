package mq

import (
	"fmt"
	"fwds/internal/conf"
)

var BusinessConfig = map[string]*Business{}

type Business struct {
	Name    string `json:"name"` //业务名称
	Topic   string `json:"topic"`
	Tag     string `json:"tag"`
	GroupId string `json:"group_id"`
}

func NewBusiness(name string, topic string, tag string, groupId string) *Business {
	prefixName := topic + tag + groupId
	if _, ok := BusinessConfig[prefixName]; ok {
		panic(fmt.Sprintf("mq key %s is exsit, please change one", prefixName))
	}
	b := &Business{Name: name, Topic: topic, Tag: tag, GroupId: groupId}
	BusinessConfig[prefixName] = b
	return b
}

func GetBusiness(key string) (*Business, error) {
	business, ok := conf.Conf.Mq.Business[key]
	if !ok {
		return nil, KeyNotFound
	}
	return &Business{
		Name:    business.Name,
		Topic:   business.Topic,
		Tag:     business.Tag,
		GroupId: business.GroupId,
	}, nil
}
