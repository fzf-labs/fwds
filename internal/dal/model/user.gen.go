// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID       int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
