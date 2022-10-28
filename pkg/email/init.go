package email

import (
	"fwds/internal/conf"
	"sync"

	"github.com/pkg/errors"
)

// Client 邮件发送客户端
var Client Driver

// Lock 读写锁
var Lock sync.RWMutex

var (
	// ErrChanNotOpen 邮件队列没有开启
	ErrChanNotOpen = errors.New("email queue does not open")
)

// Init 初始化客户端
func Init(cfg *conf.EmailConfig) {

	Lock.Lock()
	defer Lock.Unlock()

	// 确保是已经关闭的
	if Client != nil {
		Client.Close()
	}

	c := NewSMTPClient(SMTPConfig{
		Name:      cfg.Name,
		Address:   cfg.Address,
		ReplyTo:   cfg.ReplyTo,
		Host:      cfg.Host,
		Port:      cfg.Port,
		Username:  cfg.Username,
		Password:  cfg.Password,
		Keepalive: cfg.KeepAlive,
	})
	c.Init()
	Client = c

}
