package test

import (
	"fmt"
	"fwds/internal/mq/test"
	"fwds/pkg/util"
	"fwds/pkg/webhook/dingtalk"
	"time"

	"fwds/internal/dao/user_dao"
	"fwds/internal/errno"
	"fwds/internal/response"
	"fwds/pkg/debug"
	"fwds/pkg/redis"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	response.Json(c, errno.Success, "hello world")
	return
}

func Test(c *gin.Context) {
	user, err := new(user_dao.UserDao).GetUser(c, "xxxxxx@qq.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := redis.Client.WithContext(c)
	r.Get(c, "fffa")
	r.Set(c, "abc", "ssd", time.Second)
	debug.Println("bug", "bugbug", debug.WithContext(c))
	response.Json(c, errno.Success, user)
	return
}

func Test2(c *gin.Context) {
	ints := make([]int, 0)
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Second)
		ints = append(ints, i)
	}
	response.Json(c, errno.Success, nil)
	return
}

func TestMq(c *gin.Context) {
	test.Producer(util.Time.NowString())
	response.Json(c, errno.Success, nil)
	return
}
func TestDingTalk(c *gin.Context) {
	err := dingtalk.SendLink("err", "测试", "1234", "https://developers.dingtalk.com/document/robots/custom-robot-access/title-72m-8ag-pqw", "https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/5099076061/p131219.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	response.Json(c, errno.Success, nil)
	return
}
