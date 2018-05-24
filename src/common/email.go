package common

import(
	"net/smtp"
	"strings"
	"net/url"
	"encoding/base64"
	"math/rand"
	"time"
	"bytes"
	"errors"
	"regexp"
	"model"
)

type Email struct {
	user string
	password string
	host string
	port string
}

var g_email *Email
type unencryptedAuth struct {
    smtp.Auth
}

func init() {
	g_email := Email{}
	g_email.user = GetConfigValue("email", "user")
	g_email.password = GetConfigValue("email", "password")
	g_email.host = GetConfigValue("email", "host")
	g_email.port = GetConfigValue("email", "port")
}

//SSL免验证
func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    s := *server
    s.TLS = true
    return a.Auth.Start(&s)
}

func checkEmailParams() bool{
	if g_email.user == "" || g_email.password == "" || g_email.host == "" || g_email.port == "" {
		return false
	}
	return true
}
//发送邮件
func sendEmail(to, subject, body, mailType string) error {
	serverParamsValid := checkEmailParams()
	if serverParamsValid == false {
		Debug("send email get server params failed")
		return errors.New("send email get server params failed")
	}
	auth := unencryptedAuth {
		smtp.PlainAuth(
			"",
			g_email.user,
			g_email.password,
			g_email.host,
		),
	}
	var content_type string
	if mailType == "html" {
		content_type = "Content-Type: text/html;charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain;charset=UTF-8"
	}
	msg := []byte("To: "+ to + "\r\nFrom: "+ g_email.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(g_email.host+":"+g_email.port, auth, g_email.user, sendTo, msg)
	if err != nil {
		Debug("email send fail:", err.Error())
	}
	return err
}
//发送注册验证邮件
func (this *Email) SendVerify(to string) (string,error){
	code := getRandomString(6)
	verifyString := strings.Join([]string{"email=", to, "&code=", code, "&sign=regist"},"")
	result, err := AesEncrypt([]byte(verifyString))
	if err != nil {
		return code, err
	}
	encryptString := base64.StdEncoding.EncodeToString(result)

	serverHost := GetConfigValue("server", "host")
	if serverHost == "" {
		Error("send email get server host failed")
		return "", errors.New("get host failed")
	}
	linkUrl := serverHost + "?flag=" + encryptString
	subject := "有赞币答注册验证"
	body := `
		<html>
		<body>
		<h5>
		你正在注册有赞币答，点击以下链接完成注册流程
		<br/>
		<a href="
		`+ linkUrl +
		`">`+linkUrl + `</a>
		</h5>
		</body>
		</html>
		`
	err = sendEmail(to, subject, body, "html")
	if err != nil {
		return "", err
	}
	return code, nil
	//err = saveToDb(to, code, "regist", 7200)
	//return code, err
}

func (this *Email) CheckValid(email string) bool {
	reg := regexp.MustCompile("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$")
	match := reg.MatchString(email)
	return match
}
//地址验证
func (this *Email) Verify(code string) error {
	decodedVerification,err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}
	originData, err := AesDecrypt(decodedVerification)
	if err != nil {
		return err
	}
	params, perr := url.ParseQuery(string(originData))
	if perr != nil {
		return perr
	}
	paramFlag := params.Get("flag")
	paramEmail := params.Get("email")
	paramCode := params.Get("code")
	if paramFlag == "" || paramEmail == "" || paramCode == "" {
		return errors.New("验证地址不正确")
	}
	return nil
	//根据email更新用户认证状态
	//err = this.UpdateUserVerifyStatus(paramEmail, "Y")
	//return nil
}
//产生随机字符串
func getRandomString(len int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	byteStr := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
	   result = append(result, byteStr[r.Intn(bytes.Count(byteStr, nil))])
	}
	return string(result)
 }


 func saveToDb(email, code, emailType string, duration int) error{
	var emailModel model.EmailVerifyModel
	emailModel.Email = email
	emailModel.Code = code
	emailModel.MailType = emailType
	emailModel.Duration = duration
	err := GetDbInstance().Where(model.EmailVerifyModel{Email: email}).FirstOrCreate(&emailModel).Error
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
	err := sendEmail(to, subject, body, "html")
	if err == nil {
		Debug("email send successful")
	}
	Debug("email send err:", err.Error())
}