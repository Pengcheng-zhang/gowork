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
	var category []model.CategoryModel
	category = comm.GetCategory(1)
	fmt.Println(category)
	var home controller.HomeController
	m.Get("/", home.Index)
	m.Post("/regist", home.Regist)
	m.Post("/api/login", home.Login)
	m.Post("/logout", home.Logout)
	m.Get("/login", home.GetLogin)
	m.Get("/regist", home.GetRegist)

	//个人中心
	var userCenter controller.UserCenterController
	m.Group("/ucenter",func (r martini.Router)  {
		//首页
		r.Get("/:id\\d{1,5}",userCenter.Index)
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
		//发表新文章view
		r.Get("/new_article", userCenter.NewArticleView)
		//文章列表
		r.Get("/art_list", userCenter.GetArtList)
	})
	var articleController controller.ArticleController
	m.Group("/article",func(r martini.Router){
		m.Get("/:id\\d{1,5}", articleController.Detail)
		m.Post("/new", articleController.NewArticle)
		m.Post("/prise", articleController.AddPriseNum)
		m.Post("/diss", articleController.AddDissNum)
		m.Post("/delete", articleController.Delete)
		m.Post("/comment", articleController.AddReply)
		m.Post("/comment_delete", articleController.DeleteReply)
		m.Post("/comment_list", articleController.GetReplyList)
	})

	var admin controller.AdminController
	var loginMiddleWare middleware.LoginRequired
	m.Group("/admin", func (r martini.Router)  {
		r.Get("", admin.Index)
		r.Get("/tech", admin.ArticleList);
	}, loginMiddleWare.Call)

	var subCategory controller.SubCategoryController
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
			User model.UserModel
			Js []string
			Css []string
			Category []model.CategoryModel
			CurrentTab int
		}
		var data output
		r.HTML(404, "404", data)
	  })
	m.Run()
}
