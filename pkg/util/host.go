package util

import "os"

var Host = newHost()

type host struct {
}

func newHost() *host {
	return &host{}
}

//
// GetHostName
// @Description:获取主机名
// @receiver h
// @return string
//
func (h *host) GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return hostname
}
