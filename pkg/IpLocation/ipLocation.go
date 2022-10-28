package iploaction

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

// SearchLocationInfo 查询单个ip的归属地
func SearchLocationInfo(ip string) (*ip2region.IpInfo, error) {
	region, err := ip2region.New("config/ip2region/ip2region.db")
	defer region.Close()
	if err != nil {
		return nil, err
	}
	ipInfo, err := region.MemorySearch(ip)
	if err != nil {
		return nil, err
	}
	return &ipInfo, nil
}

// SearchLocation 查询单个ip的归属地
func SearchLocation(ip string) (string, error) {
	region, err := ip2region.New("config/ip2region/ip2region.db")
	defer region.Close()
	if err != nil {
		return "", err
	}
	ipInfo, err := region.MemorySearch(ip)
	if err != nil {
		return "", err
	}
	return ipInfo.Country + " " + ipInfo.Province + " " + ipInfo.City + " " + ipInfo.ISP, nil
}

// SearchBatchLocation 查询多个ip的归属地
func SearchBatchLocation(ips []string) (map[string]string, error) {
	region, err := ip2region.New("config/ip2region/ip2region.db")
	defer region.Close()
	ipList := make(map[string]string)
	if err != nil {
		return ipList, err
	}
	for _, v := range ips {
		ipInfo, err := region.MemorySearch(v)
		if err != nil {
			ipList[v] = ""
		} else {
			ipList[v] = ipInfo.Country + " " + ipInfo.Province + " " + ipInfo.City + " " + ipInfo.ISP
		}
	}
	return ipList, nil
}
