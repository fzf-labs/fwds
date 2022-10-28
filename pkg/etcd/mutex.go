package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func Mutex(client *clientv3.Client, key string, f func() error, sopts ...concurrency.SessionOption) error {
	s, err := concurrency.NewSession(client, sopts...)
	if err != nil {
		return err
	}
	defer s.Close()

	m := concurrency.NewMutex(s, key)
	if err := m.Lock(context.TODO()); err != nil {
		return err
	}
	defer func() {
		if err := m.Unlock(context.TODO()); err != nil {
			fmt.Println("etcd mux unlock err")
		}
	}()
	return f()
}
