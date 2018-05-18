package controller

import (
	"strings"
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"biz"
	"model"
	"net/http"
	"strconv"
)

type HomeController struct{
	uCenterBiz biz.UserBiz
	commBiz biz.CommomBiz
	artBiz biz.ArtBiz
	hResult htmlResult  //html数据
	jResult interface{} //api请求返回结果
}

//首页 / Get
func (this *HomeController) Index(r render.Render, session sessions.Session) {
	this.hResult.User = GetUser(session)
	this.hResult.Category = this.commBiz.GetCategory(3)
	this.hResult.CurrentCate = this.commBiz.GetCategoryByName("/tab/jobs")
	types := this.commBiz.GetSubcateIds(3)
	this.hResult.Articles = this.artBiz.GetTabArtList(types, 50, 0, "P")
	r.HTML(200, "index", this.hResult)
}

//登录 /login Post
func (this *HomeController) Login(r render.Render, req *http.Request, session sessions.Session)  {
	email := req.FormValue("email")
	password := req.FormValue("password")
	if email == "" || password == "" {
		this.jResult = map[string]interface{}{"code": 10001, "message": "请输入邮箱和密码", "result":""}
		r.JSON(200, this.jResult)
		return
	}
	var user model.UserModel
	loginSession,user, err := this.uCenterBiz.Login(email, password)
	if err != nil {
		this.jResult = map[string]interface{}{"code": 10001, "message" : "用户名或密码错误"}
		r.JSON(200, this.jResult)
		return
	}
	session.Set("yz_session_token", loginSession)
	var nextUrl string = "/"
	if strings.Index(user.Roles, "A") != -1 {
		nextUrl = "/admin"
	}
	this.jResult = map[string]interface{}{"code": 10000, "message" : "success", "result": nextUrl}
	r.JSON(200, this.jResult)
}

//登陆页 /login GET
func (this *HomeController) GetLogin(r render.Render, session sessions.Session)  {
	user := GetUser(session)
	if user.Id > 0 {
		r.Redirect("/")
	}
	this.hResult.User = user
	this.hResult.Js = []string{"/js/yzcomm.js"}
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
	var emailBiz biz.EmailBiz
	match := emailBiz.CheckValid(email)
	if ! match {
		return false, "请填写正确的邮箱"
	}
	return true,""
}
//注册 /regist POST
func (this *HomeController) Regist(r render.Render, req *http.Request, session sessions.Session) {
	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	check,message := this.checkSignupParams(username, email, password)
	if ! check {
		this.jResult = map[string]interface{}{"error": 10001, "message" : message}
		r.JSON(200, this.jResult)
		return 
	}
	var nextUrl string
	message, success := this.uCenterBiz.Register(username, email, password)
	if success {
		var user model.UserModel
		loginSession,user,err := this.uCenterBiz.Login(email, password)
		fmt.Printf("user login: login session: %s\n", loginSession)
		if err != nil {
			this.jResult = map[string]interface{}{"error": 10001, "message" : err}
			r.JSON(200, this.jResult)
			return
		}
		fmt.Printf("session=%s\n", loginSession)
		session.Set("yz_session_token", loginSession)
		nextUrl = strings.Join([]string{"/user/", strconv.Itoa(user.Id)}, "")
		this.jResult = map[string]interface{}{"code": 10000, "message" : "success", "result": nextUrl}
	} else {
		this.jResult = map[string]interface{}{"code": 10001, "message" : message}
	}
	r.JSON(200, this.jResult)
}

//登出 /api/logout POST
func (this *HomeController) Logout(r render.Render, session sessions.Session) {
	session.Set("yz_session_token", "")
	this.jResult = map[string]interface{}{"code": 10000, "message" : "success", "result": ""}
	r.JSON(200, this.jResult)
}

//关于
func (this *HomeController) About(r render.Render, session sessions.Session) {
	r.HTML(200, "main/about", this.hResult)
}