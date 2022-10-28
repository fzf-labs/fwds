package util

import (
	guuid "github.com/google/uuid"
	"github.com/teris-io/shortid"
)

var ID = newUuid()

type uuid struct {
}

func newUuid() *uuid {
	return &uuid{}
}

// GenUUID 生成随机字符串，eg: 76d27e8c-a80e-48c8-ad20-e5562e0f67e4
func (uu *uuid) GenUUID() string {
	u, _ := guuid.NewRandom()
	return u.String()
}

// GenShortID 生成一个id
func (uu *uuid) GenShortID() (string, error) {
	return shortid.Generate()
}
