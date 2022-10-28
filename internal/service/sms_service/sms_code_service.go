package sms_service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fwds/pkg/log"
	"fwds/pkg/redis"

	"github.com/pkg/errors"
)

var SmsVerifyService = NewSmsVerify()

const (
	verifyCodeRedisKey = "sms:code:%d"    // 验证码key
	maxDurationTime    = 10 * time.Minute // 验证码有效期
)

func NewSmsVerify() ISmsCodeService {
	return &SmsCodeService{}
}

type ISmsCodeService interface {
	GenSmsCode(phone int) (int, error)
	CheckSmsCode(phone, code int) bool
	GetSmsCode(phone int) (int, error)
}

type SmsCodeService struct{}

//生成短信验证码
func (s *SmsCodeService) GenSmsCode(phone int) (int, error) {
	codeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		return 0, errors.Wrap(err, "string convert int err")
	}
	key := fmt.Sprintf(verifyCodeRedisKey, phone)
	err = redis.Client.Set(context.Background(), key, code, maxDurationTime).Err()
	if err != nil {
		return 0, errors.Wrap(err, "gen login code from redis set err")
	}
	return code, nil
}

// 校验短信验证码
func (s *SmsCodeService) CheckSmsCode(phone, code int) bool {
	oldCode, err := s.GetSmsCode(phone)
	if err != nil {
		log.SugaredLogger.Warnf("[sms_code_service] get sms code err, %v", err)
		return false
	}
	if code != oldCode {
		return false
	}
	return true
}

// 获取sms的code
func (s *SmsCodeService) GetSmsCode(phone int) (int, error) {
	key := fmt.Sprintf(verifyCodeRedisKey, phone)
	code, err := redis.Client.Get(context.Background(), key).Result()
	if err == redis.ErrRedisNotFound {
		return 0, nil
	} else if err != nil {
		return 0, errors.Wrap(err, "redis get sms code err")
	}
	verifyCode, err := strconv.Atoi(code)
	if err != nil {
		return 0, errors.Wrap(err, "strconv err")
	}
	return verifyCode, nil
}
