package users_model

import "time"

// Users
//go:generate gormgen -structs Users -input .
type Users struct {
	Id            int64     //
	Username      string    // 用户名称
	Password      string    // 密码
	Nickname      string    // 昵称
	Email         string    // 邮箱
	Phone         string    // 手机
	Avatar        string    // 头像
	Bio           string    //
	Address       string    // 地址
	Job           string    // 工作
	Salt          string    //
	Uuid          string    //
	LastIp        string    // 最后一次登录IP
	LastLoginTime time.Time `gorm:"time"` // 最后一次登录时间
	Type          int32     // 用户类型
	Status        int32     //
	DeletedAt     time.Time `gorm:"time"` //
	CreatedAt     time.Time `gorm:"time"` //
	UpdatedAt     time.Time `gorm:"time"` //
}
