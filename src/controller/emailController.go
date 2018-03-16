package controller

import (
	"fmt"
	"model"
	"biz"
	"net/http"
	"github.com/martini-contrib/render"
)

type EmailController struct{
	emailBiz biz.EmailBiz
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
	verifyCode := req.FormValue("code")
	if len(verifyCode) > 0 {
		err := this.emailBiz.Verify(verifyCode)
		if err == nil {
			eResult.Code = 10000
			eResult.Message = "注册验证成功"
		}
	}
	r.HTML(200, "success", eResult)
}

func (this *EmailController) SendRegistVerification(r render.Render, req *http.Request) {
	toEmail := req.FormValue("email")
	fmt.Println("email = ", toEmail)
}