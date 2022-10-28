package sms_service

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/sms"
	"github.com/spf13/viper"
)

type ISmsSendService interface {
	Send(phone int, code int) error
	SendByQiNiu(phone int, code int) error
}

func NewSmsSendService() ISmsSendService {
	return &SmsSendService{}
}

type SmsSendService struct{}

func (s *SmsSendService) Send(phone int, code int) error {
	// 校验参数的正确性
	if phone == 0 || code == 0 {
		return errors.New("param phone or verify_code error")
	}
	return s.SendByQiNiu(phone, code)
}

//发送短信
func (s *SmsSendService) SendByQiNiu(phone int, code int) error {
	phoneStr := strconv.Itoa(phone)
	accessKey := viper.GetString("qiniu.access_key")
	secretKey := viper.GetString("qiniu.secret_key")
	mac := auth.New(accessKey, secretKey)
	manager := sms.NewManager(mac)
	args := sms.MessagesRequest{
		SignatureID: viper.GetString("qiniu.signature_id"),
		TemplateID:  viper.GetString("qiniu.template_id"),
		Mobiles:     []string{phoneStr},
		Parameters: map[string]interface{}{
			"code": code,
		},
	}

	ret, err := manager.SendMessage(args)
	if err != nil {
		return errors.Wrap(err, "send sms message error")
	}

	if len(ret.JobID) == 0 {
		return errors.New("send sms message job id error")
	}

	return nil
}
