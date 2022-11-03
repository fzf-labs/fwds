package monitor

import (
	"net/http"
	_ "net/http/pprof"
)

func NewPprofServer(addr string) error {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}
