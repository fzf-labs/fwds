package main

import (
	"flag"
	"fmt"
	"fwds/pkg/conversion"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"path"
)

var dataMap = map[string]func(detailType string) (dataType string){
	//"int":     func(detailType string) (dataType string) { return "int64" },
	//"tinyint": func(detailType string) (dataType string) { return "int32" },
	"json": func(string) string { return "datatypes.JSON" },
}
var FileName = flag.String("f", "", "the config file")

func main() {
	flag.Parse()
	dsn := GetDsn(*FileName)
	db := ConnectDB(dsn)
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/dal/query",
		ModelPkgPath: "./internal/dal/model",
	})
	g.UseDB(db)
	g.WithDataTypeMap(dataMap)
	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)
	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(model.Method) {}, g.GenerateModel("game_home"))
	g.Execute()
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

func ConnectDB(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}
	return db
}
