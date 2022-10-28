package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// BaseModel
// @Description: 1. 自定义 BaseModel，结构和 gorm.Model 一致，将 time.Time 替换为 Xtime；
//
type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt Time
	UpdatedAt Time
	DeletedAt *Time `sql:"index"`
}

// Time
// @Description:2. 创建 time.Time 类型的副本 XTime；
//
type Time struct {
	time.Time
}

//
// MarshalJSON
// @Description:3. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
// @receiver t
// @return []byte
// @return error
//
func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

//
// Value
// @Description:4. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
// @receiver t
// @return driver.Value
// @return error
//
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Predicate is a string that acts as a condition in the where clause
type Predicate string

var (
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
)
