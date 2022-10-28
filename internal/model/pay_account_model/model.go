package pay_account_model

import (
	"fwds/internal/model/pay_account_api_model"
	"time"
)

// PayAccount 支付账户
// TableName: pay_account
//go:generate gormgen -structs PayAccount -input .
type PayAccount struct {
	Id         int64     //
	Name       string    // 名称
	AppKey     string    // 应用关键词
	AppSecret  string    // 应用秘钥
	Status     int32     // 1 启用 -1 禁用 -2删除
	CreateTime time.Time `gorm:"time"` //
	UpdateTime time.Time `gorm:"time"` //

	PayApis []pay_account_api_model.PayAccountApi
}
