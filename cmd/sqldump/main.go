package main

import (
	"flag"
	"fmt"
	"fwds/pkg/conversion"
	"fwds/pkg/util"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path"
	"regexp"
	"strings"
)

var FileName = flag.String("f", "", "the config file")

// SHOW CREATE TABLE `alili_shortbook`.`shortbook_comment_reply_like`;
// SELECT table_name FROM information_schema.tables WHERE table_schema='alili_shortbook';
func main() {
	flag.Parse()
	dsn := GetDsn(*FileName)
	database := regexp.MustCompile(`/(.*?)\?`).FindString(dsn)
	database = strings.TrimLeft(database, "/")
	database = strings.TrimRight(database, "?")
	//连接数据库
	db := ConnectDB(dsn)
	var tables []string
	//查所有的table
	db.Raw(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema='%s'", database)).Pluck("table_name", &tables)
	if len(tables) == 0 {
		panic("SELECT NO TABLES")
	}
	type Result struct {
		Table       string `gorm:"column:Table"`
		CreateTable string `gorm:"column:Create Table"`
	}
	str := ""
	for _, v := range tables {
		var res Result
		db.Raw(fmt.Sprintf("SHOW CREATE TABLE %s", v)).Scan(&res)
		if res.CreateTable != "" {
			str += res.CreateTable + ";\n"
		}
	}
	target, _ := util.Str.SubstrTarget(*FileName, "/app", "left", false)
	p := target + "/storage/sql/" + database + ".sql"
	err := util.File.WriteWithIo(p, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SQL CREATE TABLE 导出成功")
}

func ConnectDB(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}
	return db
}

func GetDsn(fileName string) string {
	fileDir := path.Dir(fileName)
	filePathBase := path.Base(fileName)
	fileExt := path.Ext(fileName)
	filePrefix := filePathBase[0 : len(filePathBase)-len(fileExt)]
	config := viper.New()
	config.AddConfigPath(fileDir)    //设置读取的文件路径
	config.SetConfigName(filePrefix) //设置读取的文件名
	config.SetConfigType("yaml")     //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conversion.String(config.Get("Mysql.Default.DSN"))
}
