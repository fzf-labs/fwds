package email

import (
	"bytes"
	"fmt"
	"fwds/internal/conf"
	"html/template"
	"os"
	"time"

	"fwds/pkg/log"
)

// ActiveUserMailData 激活用户模板数据
type ActiveUserMailData struct {
	UserName      string `json:"user_name"`
	HomeURL       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ActivateURL   string `json:"activate_url"`
	Year          int    `json:"year"`
}

// NewActivationHTMLEmail 发送激活邮件 html
func NewActivationHTMLEmail(username, activateURL string) (subject string, body string) {
	mailData := ActiveUserMailData{
		UserName:      username,
		HomeURL:       conf.Conf.App.Host,
		WebsiteName:   conf.Conf.App.Name,
		WebsiteDomain: conf.Conf.App.Host,
		ActivateURL:   activateURL,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/active-mail.html", mailData)
	return "帐号激活链接", mailTplContent
}

// VerificationCodeData
// @Description: 邮件验证码
type VerificationCodeData struct {
	HomeURL       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	Code          string `json:"code"`
	Greeting      string `json:"greeting"`
	Intro         string `json:"intro"`
	Outro         string `json:"outro"`
	Year          int    `json:"year"`
}

func NewVerificationCode(code string, greeting string, intro string, outro string) (subject string, body string) {
	mailData := &VerificationCodeData{
		HomeURL:       conf.Conf.App.Host,
		WebsiteName:   conf.Conf.App.Name,
		WebsiteDomain: conf.Conf.App.Host,
		Code:          code,
		Greeting:      greeting,
		Intro:         intro,
		Outro:         outro,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/verification-code.html", mailData)
	return "验证码", mailTplContent
}

// ResetPasswordMailData 激活用户模板数据
type ResetPasswordMailData struct {
	UserName      string `json:"user_name"`
	HomeURL       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ResetURL      string `json:"reset_url"`
	Year          int    `json:"year"`
}

// NewResetPasswordHTMLEmail 发送重置密码邮件 html
func NewResetPasswordHTMLEmail(username, resetURL string) (subject string, body string) {
	mailData := ResetPasswordMailData{
		UserName:      username,
		HomeURL:       conf.Conf.App.Host,
		WebsiteName:   conf.Conf.App.Name,
		WebsiteDomain: conf.Conf.App.Host,
		ResetURL:      resetURL,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/reset-mail.html", mailData)
	return "密码重置", mailTplContent
}

type NotifyMailData struct {
	UserName      string `json:"user_name"`
	HomeURL       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	JumpURL       string `json:"jump_url"`
	Greeting      string `json:"greeting"`
	Intro         string `json:"intro"`
	Outro         string `json:"outro"`
	Year          int    `json:"year"`
}

func NewNotifyMailData(jumpURL string, greeting string, intro string, outro string) (subject string, body string) {
	mailData := &NotifyMailData{
		JumpURL:  jumpURL,
		Greeting: greeting,
		Intro:    intro,
		Outro:    outro,
		Year:     time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/notify-mail.html", mailData)
	return "告警", mailTplContent
}

// getEmailHTMLContent 获取邮件模板
func getEmailHTMLContent(tplPath string, mailData interface{}) string {
	b, err := os.ReadFile(tplPath)
	if err != nil {
		log.SugaredLogger.Warnf("[util.email] read file err: %v", err)
		return ""
	}
	mailTpl := string(b)
	tpl, err := template.New("email tpl").Parse(mailTpl)
	if err != nil {
		log.SugaredLogger.Warnf("[util.email] template new err: %v", err)
		return ""
	}
	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, mailData)
	if err != nil {
		fmt.Println("exec err", err)
		log.SugaredLogger.Warnf("[util.email] execute template err: %v", err)
	}
	return buffer.String()
}
