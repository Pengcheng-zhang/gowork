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

type htmlResult struct {
	Js []string
	Css []string
	CurrentTab int
	Data interface{}
}

var uCenter biz.UCenterManager
var comm biz.Commom
var hResult htmlResult  //html数据
var jResult interface{} //api请求返回结果

//首页 / Get
func (this *HomeController) Index(r render.Render, session sessions.Session) {
	v := session.Get("sucai_session_token")
	fmt.Println(v)
	var user model.UserModel
	if v != nil {
		user = biz.GetUserFromSession(v.(string))
		fmt.Println(user)
	}
	hResult.Data = map[string]interface{}{"User": user, "Category": comm.GetCategory(1)}
	hResult.CurrentTab = 1
	fmt.Println(hResult)
	r.HTML(200, "index", hResult)
	//r.JSON(200, htmlResult)
}

//登录 /login Post
func (this *HomeController) Login(r render.Render, req *http.Request, session sessions.Session)  {
	email := req.FormValue("email")
	password := req.FormValue("password")
	if email == "" || password == "" {
		jResult = map[string]interface{}{"code": 10001, "message": "请输入邮箱和密码", "result":""}
		r.JSON(200, jResult)
		return
	}
	var user model.UserModel
	loginSession,user, err := uCenter.Login(email, password)
	if err != nil {
		jResult = map[string]interface{}{"code": 10001, "message" : err}
		r.JSON(200, jResult)
		return
	}
	session.Set("sucai_session_token", loginSession)
	var nextUrl string = "/"
	if strings.Index(user.Roles, "A") != -1 {
		nextUrl = "/admin"
	}
	jResult = map[string]interface{}{"error": 10000, "message" : "success", "next_url": nextUrl}
	r.JSON(200, jResult)
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
	hResult.Data = map[string]interface{}{"User": user}
	hResult.Js = []string{"/js/yzcomm.js"}
	r.HTML(200, "main/signin", hResult)
}

func (this *HomeController) GetRegist(r render.Render, session sessions.Session)  {
	r.HTML(200, "main/signup", hResult)
}

//注册 /regist POST
func (this *HomeController) Regist(r render.Render, req *http.Request, session sessions.Session) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	fmt.Printf("email:%s\tpassword:%s\n", email, password)
	if email == "" || password == "" {
		jResult = map[string]interface{}{"error": 10001, "message" : "请输入邮箱和密码"}
		r.JSON(200, jResult)
		return
	}
	var success bool
	var nextUrl string
	success = uCenter.Register(email, password)
	if success {
		var user model.UserModel
		loginSession,user,err := uCenter.Login(email, password)
		fmt.Printf("user login: login session: %s\n", loginSession)
		if err != nil {
			jResult = map[string]interface{}{"error": 10001, "message" : err}
			r.JSON(200, jResult)
			return
		}
		fmt.Printf("session=%s\n", loginSession)
		session.Set("sucai_session_token", loginSession)
		nextUrl = strings.Join([]string{"/user/", strconv.Itoa(user.Id)}, "")
	}
	jResult = map[string]interface{}{"error": 10000, "message" : "success", "next_url": nextUrl}
	r.JSON(200, jResult)
}

//登出 /api/logout POST
func (this *HomeController) Logout(r render.Render, session sessions.Session) {
	session.Set("sucai_session_token", "")
	r.Redirect("/")
}