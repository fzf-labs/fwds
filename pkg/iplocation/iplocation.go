package iplocation

import (
	"fmt"
	"fwds/pkg/once"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var oo once.Once

var il ipLocation

type ipLocation struct {
	searcher *xdb.Searcher
}

func NewIpLocation(dbPath string) (*ipLocation, error) {
	err := oo.Do(func() error {
		// 1、从 dbPath 加载整个 xdb 到内存
		cBuff, err := xdb.LoadContentFromFile(dbPath)
		if err != nil {
			fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
			return err
		}

		// 2、用全局的 cBuff 创建完全基于内存的查询对象。
		searcher, err := xdb.NewWithBuffer(cBuff)
		if err != nil {
			fmt.Printf("failed to create searcher with content: %s\n", err)
			return err
		}
		il.searcher = searcher
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &il, nil
}

// SearchLocation 查询单个ip的归属地
func (i *ipLocation) SearchLocation(ip string) (string, error) {
	str, err := i.searcher.SearchByStr(ip)
	if err != nil {
		return "", err
	}
	return str, nil
}

// SearchBatchLocation 查询多个ip的归属地
func (i *ipLocation) SearchBatchLocation(ips []string) (map[string]string, error) {
	items := make(map[string]string)
	for _, ip := range ips {
		str, err := i.searcher.SearchByStr(ip)
		if err != nil {
			return nil, err
		}
		items[ip] = str
	}
	return items, nil
}
