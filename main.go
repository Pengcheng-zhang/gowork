package main

import (
	"fmt"
	"model"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"biz"
	"github.com/martini-contrib/sessions"
	"controller"
	"middleware"
)

func main()  {
	biz.DbInit();
	defer biz.GetDbInstance().Close()
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:"templates",
		Layout: "layout",
		Extensions:[]string{".tmpl",".html"},
		Charset:"UTF-8",
		IndentJSON: true,
	}))
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	
	var comm biz.Commom
	var category []model.Category
	category = comm.GetCategory(1)
	fmt.Println(category)
	var home controller.Home
	m.Get("/", home.Index)
	m.Post("/regist", home.Regist)
	m.Post("/api/login", home.Login)
	m.Post("/logout", home.Logout)
	m.Get("/login", home.GetLogin)
	m.Get("/regist", home.GetRegist);

	//个人中心
	var userCenter controller.UserCenter
	m.Group("/ucenter",func (r martini.Router)  {
		//首页
		r.Get("/:id\\d+",userCenter.Index)
		//收藏管理
		r.Get("/collect", userCenter.Collections)
		//签到记录
		r.Get("/sign_history", userCenter.SignLogHistory)
		//我的积分
		r.Get("/points", userCenter.MyPoints)
		//站内信息
		r.Get("/message",userCenter.Message)
		//邀请好友
		r.Get("/invite", userCenter.InviteFriends)
		//个人资料
		r.Get("/info", userCenter.SelfInfo)
		//修改密码
		r.Get("/password", userCenter.ChangePassword)
		//发表新文章
		r.Get("/new_article", userCenter.NewArticle)
	})

	var admin controller.Admin
	var loginMiddleWare middleware.LoginRequired
	m.Group("/admin", func (r martini.Router)  {
		r.Get("", admin.Index)
		r.Get("/tech", admin.ArticleList);
	}, loginMiddleWare.Call)

	var subCategory controller.SubCategory
	m.Group("/go", func (r martini.Router)  {
		r.Get("/php", subCategory.ToPHP)
		r.Get("/python", subCategory.ToPython)
		r.Get("/java", subCategory.ToJava)
		r.Get("/nodejs", subCategory.ToNodeJs)
		r.Get("/golang", subCategory.ToGoLang)
		r.Get("/android", subCategory.ToAndroid)
		r.Get("/ios", subCategory.ToIOS)
	})

	m.NotFound(func(r render.Render) {
		type output struct {
			User model.User
			Js []string
			Css []string
			Category []model.Category
			CurrentTab int
		}
		var data output
		r.HTML(404, "404", data)
	  })
	m.Run()
}
