package config

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"strings"
	"time"
)

func InitByEtcd() {
	viperFile, err := LoadConfig("")
	if err != nil {
		return
	}
	host := viperFile.GetString("Etcd.Host")
	configPath := viperFile.GetString("Etcd.ConfigPath")
	secretKey := viperFile.GetString("Etcd.SecretKey")
	viperEtcd := viper.New()
	//https
	if strings.HasPrefix(host, "https://") {
		err := viperEtcd.AddSecureRemoteProvider("etcd", host, configPath, secretKey)
		if err != nil {
			log.Fatal("config init by etcd https err")
			return
		}
	} else {
		err := viperEtcd.AddRemoteProvider("etcd", host, configPath)
		if err != nil {
			log.Fatal("config init by etcd http err")
			return
		}
	}
	viperEtcd.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(configPath)), ".")) // Need to explicitly set this to yaml
	err = viperEtcd.ReadRemoteConfig()
	if err != nil {
		log.Fatalf("config read by etcd err: %v", err)
	}
	err = ParseConfig(viperEtcd)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	go func() {
		for {
			time.Sleep(time.Second * 5) // 每次请求后延迟一下
			// 目前只测试了etcd支持
			err := viperEtcd.WatchRemoteConfig()
			if err != nil {
				log.Printf("unable to read remote config: %v", err)
				continue
			}
			err = ParseConfig(viperEtcd)
			if err != nil {
				log.Printf("etcd watch config err:%v", err)
			}
		}
	}()

	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}
