package dingtalk

import (
	"encoding/json"
	"fmt"
	"fwds/internal/conf"
	"fwds/pkg/bcrypt"
	"fwds/pkg/httpclient"
	"fwds/pkg/util"
	"github.com/pkg/errors"
)

type reqText struct {
	At struct {
		AtMobiles []string `json:"atMobiles"` //被@人的手机号。
		IsAtAll   bool     `json:"isAtAll"`   //是否@所有人。
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	MsgType string `json:"msgtype"` //消息类型，此时固定为：text。
}

type reqLink struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
}

func SendText(chat, msg string, at ...string) error {
	IsAtAll := true
	if len(at) > 0 {
		IsAtAll = false
	}
	req := reqText{
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{
			AtMobiles: at,
			IsAtAll:   IsAtAll,
		},
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
		MsgType: "text",
	}
	return send(chat, req)
}

func SendLink(chat, title, text, messageUrl, picUrl string) error {
	req := reqLink{
		Msgtype: "link",
		Link: struct {
			Text       string `json:"text"`
			Title      string `json:"title"`
			PicUrl     string `json:"picUrl"`
			MessageUrl string `json:"messageUrl"`
		}{
			Text:       text,
			Title:      title,
			MessageUrl: messageUrl,
			PicUrl:     picUrl,
		},
	}
	return send(chat, req)
}

//发送消息
func send(chat string, req interface{}) error {
	business, ok := conf.Conf.DingTalk[chat]
	if !ok {
		return errors.New("未配置的钉钉群")
	}
	jsonReq, _ := json.Marshal(req)
	timestamp := util.Time.Now().TimestampWithMillisecond()
	sign := bcrypt.HmacSha256(fmt.Sprintf("%d\n%s", timestamp, business.Secret), business.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", business.Url, timestamp, sign)
	_, err := httpclient.PostJSON(url, jsonReq)
	if err != nil {
		return err
	}
	return nil
}
