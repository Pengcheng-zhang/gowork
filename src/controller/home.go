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

}

var uCenter biz.UCenterManager
var comm biz.Commom
//首页 / Get
func (this *HomeController) Index(r render.Render, session sessions.Session) {
	v := session.Get("sucai_session_token")
	fmt.Println(v)
	var user model.UserModel
	if v != nil {
		user = biz.GetUserFromSession(v.(string))
		fmt.Println(user)
	}

	type output struct {
		User model.UserModel
		Js []string
		Css []string
		Category []model.CategoryModel
		CurrentTab int
	}
	var data output
	data.User = user
	data.Category = comm.GetCategory(1)
	data.CurrentTab = 1
	data.Js = []string{}
	data.Css = []string{}
	fmt.Println(data)
	r.HTML(200, "index", data)
}

//登录 /login Post
func (this *HomeController) Login(r render.Render, req *http.Request, session sessions.Session)  {
	email := req.FormValue("email")
	password := req.FormValue("password")
	if email == "" || password == "" {
		r.JSON(200, map[string]interface{}{"code": 10001, "message" : "请输入邮箱和密码"})
	}
	var user model.UserModel
	loginSession,user, err := uCenter.Login(email, password)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10001, "message" : err})
	}
	session.Set("sucai_session_token", loginSession)
	var nextUrl string = "/"
	if strings.Index(user.Roles, "A") != -1 {
		nextUrl = "/admin"
	}
	r.JSON(200, map[string]interface{}{"error": 10000, "message" : "success", "next_url": nextUrl})
}
//登陆页 /login GET
func (this *HomeController) GetLogin(r render.Render, session sessions.Session)  {
	v := session.Get("sucai_session_token")
	var user model.UserModel
	if v != nil {
		fmt.Println(v)
		user = biz.GetUserFromSession(v.(string))
		if user.Id > 0 {
			r.Redirect("/")
		}
	}
	type output struct {
		User model.UserModel
		Js []string
		Css []string
	}
	var data output
	data.User = user
	data.Js = []string{"/js/yzcomm.js"}
	data.Css = []string{}
	r.HTML(200, "login", data)
}
func (this *HomeController) GetRegist(r render.Render, session sessions.Session)  {
	
}
//注册 /regist POST
func (this *HomeController) Regist(r render.Render, req *http.Request, session sessions.Session) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	fmt.Printf("email:%s\tpassword:%s\n", email, password)
	if email == "" || password == "" {
		r.JSON(200, map[string]interface{}{"error": 10001, "message" : "请输入邮箱和密码"})
	}
	var success bool
	var nextUrl string
	success = uCenter.Register(email, password)
	fmt.Printf("user Register: success: %s\n", success)
	if success {
		var user model.UserModel
		loginSession,user,err := uCenter.Login(email, password)
		fmt.Printf("user login: login session: %s\n", loginSession)
		if err != nil {
			r.JSON(200, map[string]interface{}{"error": 10001, "message" : err})
		}
		fmt.Printf("session=%s\n", loginSession)
		session.Set("sucai_session_token", loginSession)
		nextUrl = strings.Join([]string{"/user/", strconv.Itoa(user.Id)}, "")
	}
	r.JSON(200, map[string]interface{}{"error": 10000, "message" : "success", "next_url": nextUrl})
}

func (this *HomeController) Logout(r render.Render, session sessions.Session) {
	session.Set("sucai_session_token", "")
	r.Redirect("/")
}