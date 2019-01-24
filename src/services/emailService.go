package services

import(
	"net/smtp"
	"strings"
	"errors"
)

type EmailService struct {
	g_mail *Email
}

type Email struct {
	user string
	password string
	host string
	port string
}

type unencryptedAuth struct {
    smtp.Auth
}

func (this *EmailService) initEmailService() error{
	emailSection := GetConfigSection("email")
	if emailSection != nil {
		this.g_mail.user = emailSection["user"]
		this.g_mail.password = emailSection["password"]
		this.g_mail.host = emailSection["host"]
		this.g_mail.port = emailSection["port"]
	}
	result := this.checkEmailParams()
	if result == true {
		return nil
	}
	return errors.New("失败")
}

//SSL免验证
func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    s := *server
    s.TLS = true
    return a.Auth.Start(&s)
}

func (this *EmailService) checkEmailParams() bool{
	if this.g_mail.user == "" || this.g_mail.password == "" || this.g_mail.host == "" || this.g_mail.port == "" {
		return false
	}
	return true
}

func (this *EmailService) getMailInstance() error{
	if this.g_mail == nil {
		err := this.initEmailService()
		return err
	}
	return nil
}
//发送邮件
func SendEmail(to, subject, body, mailType string) error {
	service := &EmailService{}
	err := service.getMailInstance()
	if err != nil {
		return errors.New("send email get server params failed")
	}
	serverParamsValid := service.checkEmailParams()
	if serverParamsValid == false {
		Debug("send email get server params failed")
		return errors.New("send email get server params failed")
	}
	auth := unencryptedAuth {
		smtp.PlainAuth(
			"",
			service.g_mail.user,
			service.g_mail.password,
			service.g_mail.host,
		),
	}
	var content_type string
	if mailType == "html" {
		content_type = "Content-Type: text/html;charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain;charset=UTF-8"
	}
	msg := []byte("To: "+ to + "\r\nFrom: "+ service.g_mail.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err = smtp.SendMail(service.g_mail.host + ":" + service.g_mail.port, auth, service.g_mail.user, sendTo, msg)
	if err != nil {
		Debug("email send fail:", string(msg), err.Error())
	}
	return err
}

 //测试发送邮件
func TestMailSend() {
	subject := "有赞币答注册验证"

	body := `
		<html>
		<body>
		<h5>
		你正在注册有赞币答，点击以下链接完成注册流程
		<br/>
		<a href="http://localhost:3000/email/verify?code=">http://localhost:3000/email/verify?code=</a>
		</h5>
		</body>
		</html>
		`
	to := "770651352@qq.com"
	err := SendEmail(to, subject, body, "html")
	if err == nil {
		Debug("email send successful")
	}
	Debug("email send err:", err.Error())
}