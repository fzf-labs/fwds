package pay_account_api_model

import "time"

// PayAccountApi 账户拥有的api
// TableName: pay_account_api
//go:generate gormgen -structs PayAccountApi -input .
type PayAccountApi struct {
	Id           int64     //
	PayAccountId int64     //
	Method       string    // 请求方式
	Api          string    // 请求地址
	Status       int32     // 状态 1:启用  -1:禁用
	CreateTime   time.Time `gorm:"time"` //
	UpdateTime   time.Time `gorm:"time"` //
}
