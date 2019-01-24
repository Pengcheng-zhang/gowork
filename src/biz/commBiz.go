package biz

import (
	"services"
	"model"
	"strings"
	"encoding/base64"
	"math/rand"
	"net/url"
	"time"
	"errors"
	"regexp"
)

type CommomBiz struct{

}

//邮箱校验
func (this *CommomBiz) CheckValid(email string) bool {
	reg := regexp.MustCompile("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$")
	match := reg.MatchString(email)
	return match
}
//产生随机字符串
func getRandomString(length int) string{
	byteStr := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
	   result[i] = byteStr[r.Intn(len(byteStr))]
	}
	return string(result)
 }

 func getEmailLinkUrl(to, code, mailType string) (string,error){
	verifyString := strings.Join([]string{"email=", to, "&code=", code, "&type=", mailType},"")
	result, err := services.AesEncrypt([]byte(verifyString))
	if err != nil {
		return code, err
	}
	encryptString := base64.StdEncoding.EncodeToString(result)

	serverHost := services.GetConfigValue("server", "host")
	if serverHost == "" {
		Error("send email get server host failed")
		return "", errors.New("get host failed")
	}
	linkUrl := serverHost + "/email/verify?flag=" + encryptString
	return linkUrl,nil
 }
//发送注册验证邮件
func (this *CommomBiz) SendRegistVerifyEmail(to string) error{
	code := getRandomString(6)
	linkUrl, err := getEmailLinkUrl(to, code, "regist")
	if err != nil {
		return  err
	}
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
	err = services.SendEmail(to, subject, body, "html")
	if err != nil {
		return err
	}
	err = saveToDb(to, code, "regist", 7200)
	return err
}

//忘记密码验证邮件发送
func (this *CommomBiz) SendForgetPasswordEmail(to string) error{
	code := getRandomString(6)
	serverHost := services.GetConfigValue("server", "host")
	linkUrl := serverHost + "/signin"
	subject := "有赞币答找回密码"
	body := `
		<html>
		<body>
		<h5>
		已设置您的初始密码为：` + code +
		`<br/>点击以下链接登录，然后重置密码。
		<a href="
		`+ linkUrl +
		`">`+linkUrl + `</a>
		</h5>
		</body>
		</html>
		`	
	err := services.SendEmail(to, subject, body, "html")
	if err != nil {
		return err
	}
	err = saveToDb(to, code, "forget", 10800)
	return err
}

func saveToDb(email, code, emailType string, duration int) error{
	var emailModel model.EmailVerifyModel
	emailModel.Email = email
	emailModel.MailType = emailType
	duration = duration + (int)(time.Now().Unix())
	db := GetDbInstance()
	err := db.Where(model.EmailVerifyModel{Email: email, MailType: emailType}).Assign(model.EmailVerifyModel{Code: code, Duration: duration}).FirstOrCreate(&emailModel).Error
	if err != nil {
		Error("update verify email failed1")
	}
	if emailType == "forget" {
		err := db.Table("yz_user").Where("email = ?", email).Update("password", code).Error
		if err != nil {
			Error("update user info failed:", err.Error())
		}
	}
	return err
 }

 func (this *CommomBiz) CheckLatestSendEmailTime(email, emailType string) bool{
	var emailModel model.EmailVerifyModel
	GetDbInstance().First(&emailModel, "email = ? AND mail_type = ?", email, emailType)
	currentTimestamp := (int)(time.Now().Unix())
	if emailModel.Duration > currentTimestamp {
		return  false
	}
	return true
 }
 //地址验证
func (this *CommomBiz) Verify(code string) (err error) {
	decodedVerification,err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return errors.New("链接有误,请重新发送邮件验证")
	}
	originData, err := services.AesDecrypt(decodedVerification)
	if err != nil {
		return errors.New("链接有误,请重新发送邮件验证")
	}
	params, err := url.ParseQuery(string(originData))
	if err != nil {
		return errors.New("链接有误,请重新发送邮件验证")
	}
	paramEmail := params.Get("email")
	paramCode := params.Get("code")
	paramType := params.Get("type")
	Debug("verify result:", paramEmail, paramCode, paramType)
	if paramEmail == "" || paramCode == "" || paramType == "" {
		return errors.New("验证地址不正确")
	}

	var emailModel model.EmailVerifyModel
	err = GetDbInstance().First(&emailModel, "email = ? AND mail_type = ? AND code = ?", paramEmail, paramType, paramCode).Error
	if err != nil {
		return errors.New("链接有误,请重新发送邮件验证")
	}
	return nil
}