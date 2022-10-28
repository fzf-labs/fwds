package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"fwds/internal/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func InitByFile(env string) error {
	vf, err := LoadConfig(env)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
		return err
	}
	err = ParseConfig(vf)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
		return err
	}
	WatchConfig(vf)
	return nil
}

// LoadConfig  使用viper加载指定文件夹下的配置文件
func LoadConfig(env string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	if env != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		v.SetConfigFile("config." + env)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		v.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	v.SetConfigType("yaml")
	// .号替换为_
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// ParseConfig 使用viper解析配置文件到全局结构体中
func ParseConfig(v *viper.Viper) error {
	err := v.Unmarshal(conf.Conf)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return err
	}
	return nil
}

// WatchConfig  监控配置文件变化并热加载程序
func WatchConfig(v *viper.Viper) {
	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(fe fsnotify.Event) {
			//viper配置发生变化了 执行响应的操作
			fmt.Println("Config file changed:", fe.Name)
			before, _ := json.Marshal(conf.Conf)
			err := v.Unmarshal(&conf.Conf)
			if err != nil {
				log.Printf("unable to decode into struct, %v", err)
				panic("Config file changed err" + err.Error())
			}
			after, _ := json.Marshal(conf.Conf)
			//对比
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(before), string(after), false)
			fmt.Println("Config change diff ", dmp.DiffPrettyText(diffs))
		})
	}()
}
