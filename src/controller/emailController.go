package controller

import (
	"model"
	"biz"
	"net/http"
	"github.com/martini-contrib/render"
	"strings"
)

type EmailController struct{
	commBiz biz.CommomBiz
}

type emailResult struct {
	User model.UserModel
	Code int
	Message string
}

func (this *EmailController) Verification(r render.Render, req *http.Request) {
	var eResult emailResult
	eResult.Code = 10001
	eResult.Message = "注册验证失败，请重新验证或联系管理员admin@youzanbida.com。"
	verifyCode := req.FormValue("flag")
	verifyCode = strings.Replace(verifyCode, " ", "+", -1)
	Debug("verify code:", verifyCode)
	if len(verifyCode) > 0 {
		err := this.commBiz.Verify(verifyCode)
		if err == nil {
			r.Redirect("/signin")
			return
		}
	}
	r.HTML(200, "success", eResult)
}

func (this *EmailController) SendRegistVerification(r render.Render, req *http.Request) {
	toEmail := req.FormValue("email")
	Debug("email = ", toEmail)
}