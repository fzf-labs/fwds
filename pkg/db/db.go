package db

import (
	"database/sql"
	"fwds/internal/conf"
	"gorm.io/gorm/schema"

	"fwds/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromopentracing "gorm.io/plugin/opentracing"
)

// DB 数据库全局变量
var DB map[string]*gorm.DB

// Init 初始化数据库
func Init(cfg conf.MysqlConfigs) {
	DB = make(map[string]*gorm.DB)
	for k, v := range cfg {
		mysqlConfig := v
		DB[k] = NewMySQL(&mysqlConfig)
	}
}

// GetDB 获取指定库名的连接
func GetDB(db ...string) *gorm.DB {
	var key string
	if len(db) > 0 {
		key = db[0]
	} else {
		key = "default"
	}
	g, ok := DB[key]
	if !ok {
		log.SugaredLogger.Panicf("get mysql db failed. db num is: %s", key)
	}
	return g
}

func GetReadDB() *gorm.DB {
	g, ok := DB["read"]
	if !ok {
		log.SugaredLogger.Panicf("get mysql db failed. db num is: %s", "read")
	}
	return g
}

func GetWriteDB() *gorm.DB {
	g, ok := DB["write"]
	if !ok {
		log.SugaredLogger.Panicf("get mysql db failed. db num is: %s", "write")
	}
	return g
}

func NewMySQL(cfg *conf.MysqlConfig) *gorm.DB {
	sqlDB, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.SugaredLogger.Panicf("open mysql failed,err: %+v", err)
	}
	// set for db connection
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.SugaredLogger.Info(cfg.DSN)
		log.SugaredLogger.Panicf("database connection failed, err: %+v", err)
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	// set trace
	err = db.Use(gromopentracing.New())
	if err != nil {
		log.SugaredLogger.Panicf("using gorm opentracing, err: %+v", err)
	}
	err = db.Use(&TracePlugin{})
	if err != nil {
		log.SugaredLogger.Panicf("using Plugin fail, err: %+v", err)
	}
	return db
}
