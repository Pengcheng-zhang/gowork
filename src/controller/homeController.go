package controller

import (
	"strings"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"biz"
	"model"
	"net/http"
	"strconv"
)

type HomeController struct{
	uCenterBiz biz.UserBiz
	categoryBiz biz.CategoryBiz
	commBiz biz.CommomBiz
	artBiz biz.ArtBiz
	hResult htmlResult  //html数据
}

//首页 / Get
func (this *HomeController) Index(r render.Render, req *http.Request, session sessions.Session) {
	in_page := req.FormValue("p")
	this.hResult.User = GetUser(session)
	// types := this.categoryBiz.GetSubcateIds(3)
	var page int = 0
	if in_page == "" {
		page = 0
	} else {
		var err error
		page,err = strconv.Atoi(in_page)
		if err != nil {
			page = 0
		}else {
			if page > 1 {
				page = page - 1
			}else {
				page = 0
			}
		}
	}
	offset := page * 50
	this.hResult.Articles, this.hResult.ArticleCount = this.artBiz.GetArtList(0, 50, offset, "P", "")
	this.hResult.ArticlePageCount = this.hResult.ArticleCount / 50 + 1
	this.hResult.ArticleCurrentPage = page
	r.HTML(200, "index", this.hResult)
}

//登录 /login Post
func (this *HomeController) ApiLogin(r render.Render, req *http.Request, session sessions.Session)  {
	email := req.FormValue("email")
	password := req.FormValue("password")
	if email == "" || password == "" {
		setJsonResult(10001, "请输入邮箱和密码", "")
		r.JSON(200, jsonResult)
		return
	}
	var user model.UserModel
	loginSession,user, err := this.uCenterBiz.Login(email, password)
	if err != nil {
		setJsonResult(10001, "用户名或密码错误", "")
		r.JSON(200, jsonResult)
		return
	}
	session.Set("yz_session_token", loginSession)
	var nextUrl string = "/"
	if strings.Index(user.Roles, "A") != -1 {
		nextUrl = "/admin"
	}
	setJsonResult(10000, "success", nextUrl)
	r.JSON(200, jsonResult)
}

//登陆页 /login GET
func (this *HomeController) GetLogin(r render.Render, session sessions.Session)  {
	user := GetUser(session)
	if user.Id > 0 {
		r.Redirect("/")
	}
	this.hResult.User = user
	r.HTML(200, "main/signin", this.hResult)
}

func (this *HomeController) GetRegist(r render.Render, session sessions.Session)  {
	r.HTML(200, "main/signup", this.hResult)
}

//注册信息检测
func (this *HomeController) checkSignupParams(username, email, password string) (bool, string){
	if username == "" || email == "" || password == "" {
		return false, "请填写完整信息"
	}
	if len(username) > 20 || len(username) < 5 {
		return false, "用户名不符合要求"
	}
	if len(password) > 20 || len(password) < 5 {
		return false, "密码不符合要求"
	}
	match := this.commBiz.CheckValid(email)
	if ! match {
		return false, "请填写正确的邮箱"
	}
	return true,""
}
//注册 /regist POST
func (this *HomeController) ApiRegist(r render.Render, req *http.Request, session sessions.Session) {
	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	check,message := this.checkSignupParams(username, email, password)
	if ! check {
		setJsonResult(10001, message, "")
		r.JSON(200, jsonResult)
		return 
	}
	var nextUrl string
	message, success := this.uCenterBiz.Register(username, email, password)
	if success {
		var user model.UserModel
		loginSession,user,err := this.uCenterBiz.Login(email, password)
		Debug("user login: login session:", loginSession)
		if err != nil {
			setJsonResult(10001, "登录失败", "")
			r.JSON(200, jsonResult)
			return
		}
		Debug("session=", loginSession)
		session.Set("yz_session_token", loginSession)
		nextUrl = strings.Join([]string{"/user/", strconv.Itoa(user.Id)}, "")
		setJsonResult(10000, "success", nextUrl)
	} else {
		setJsonResult(10001, message, "")
	}
	r.JSON(200, jsonResult)
}

//登出 /api/logout POST
func (this *HomeController) ApiLogout(r render.Render, session sessions.Session) {
	session.Set("yz_session_token", "")
	setJsonResult(10000, "success", "")
	r.JSON(200, jsonResult)
}

//关于
func (this *HomeController) About(r render.Render, session sessions.Session) {
	r.HTML(200, "main/about", this.hResult)
}

//忘记密码
func (this *HomeController) Forget(r render.Render, session sessions.Session) {
	r.HTML(200, "main/forget", this.hResult)
}

func (this *HomeController) ApiForgetPassword(r render.Render, req *http.Request, session sessions.Session) {
	email := req.FormValue("email")
	result := this.commBiz.CheckValid(email)
	if result == false {
		setJsonResult(10001, "邮箱地址不正确", "")
	} else {
		result = this.commBiz.CheckLatestSendEmailTime(email, "forget")
		if result == false {
			setJsonResult(10003, "操作过快，请稍后再试。", "")
		}else {
			err := this.commBiz.SendForgetPasswordEmail(email)
			if err == nil {
				setJsonResult(10000, "邮件发送成功，请前往邮箱完成下一步操作。", "")
			}else {
				setJsonResult(10002, "邮件发送失败，请稍后再试！", "")
			}
		}
	}
	r.JSON(200, jsonResult)
}