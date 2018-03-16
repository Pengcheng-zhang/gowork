package biz

import(
	"net/smtp"
	"strings"
	"fmt"
	"net/url"
	"encoding/base64"
	"math/rand"
	"time"
	"bytes"
	"errors"
	"regexp"
	"model"
)

type EmailBiz struct {

}

var _email_user string = "admin@youzanbida.com"
var _email_password string = "youzanbida123$"
var _email_host string = "smtp.mxhichina.com"
var _email_port string = "25"

type unencryptedAuth struct {
    smtp.Auth
}

//SSL免验证
func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    s := *server
    s.TLS = true
    return a.Auth.Start(&s)
}

//发送邮件
func sendEmail(to, subject, body, mailType string) error {
	auth := unencryptedAuth {
		smtp.PlainAuth(
			"",
			_email_user,
			_email_password,
			_email_host,
		),
	}
	var content_type string
	if mailType == "html" {
		content_type = "Content-Type: text/html;charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain;charset=UTF-8"
	}
	msg := []byte("To: "+ to + "\r\nFrom: "+ _email_user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(_email_host+":"+_email_port, auth, _email_user, sendTo, msg)
	if err != nil {
		fmt.Println("email send fail:", err)
	}
	return err
}
//发送注册验证邮件
func (this *EmailBiz) SendVerify(to string) (string,error){
	code := getRandomString(6)
	verifyString := strings.Join([]string{"email=", to, "&code=", code, "&sign=regist"},"")
	result, err := AesEncrypt([]byte(verifyString))
	if err != nil {
		return code, err
	}
	encryptString := base64.StdEncoding.EncodeToString(result)

	linkUrl := GetHost() + "?flag=" + encryptString
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
	err = saveToDb(to, code, "regist", 7200)
	return code, err
}

func (this *EmailBiz) CheckValid(email string) bool {
	reg := regexp.MustCompile("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$")
	match := reg.MatchString(email)
	return match
}
//地址验证
func (this *EmailBiz) Verify(code string) error {
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
	//根据email更新用户认证状态
	err = this.UpdateUserVerifyStatus(paramEmail, "Y")
	return nil
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
 func (this *EmailBiz) UpdateUserVerifyStatus(email,status string) error{
	 var userBiz UserBiz
	 var user model.UserModel
	 user.Email = email
	updateData := map[string]interface{}{"verified": status}
	_,err := userBiz.UpdateUserInfo(user, updateData)
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
		fmt.Println("email send successful")
	}
	fmt.Println("email send err:", err)
}