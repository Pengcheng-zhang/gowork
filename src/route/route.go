package route

import(
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"biz"
	"controller"
	"middleware"
	"model"
)

func Run(m *martini.ClassicMartini)  {
	var comm biz.Commom
	var category []model.CategoryModel
	category = comm.GetCategory(1)
	fmt.Println(category)
	
	commRoute(m)
	adminRoute(m)
	ucenterRoute(m)
	articleRoute(m)
	subPathRoute(m)
	route404(m)

	m.Run()
}

//通用接口路由
func commRoute(m *martini.ClassicMartini)  {
	var home controller.HomeController
	m.Get("/", home.Index) 					//首页
	m.Get("/login", home.GetLogin)			//登陆页
	m.Get("/regist", home.GetRegist)		//注册页

	m.Post("/api/regist", home.Regist)  		//注册
	m.Post("/api/login", home.Login)		//登陆
	m.Post("/api/logout", home.Logout)			//登出
}

//后台管理路由
func adminRoute(m *martini.ClassicMartini)  {
	var admin controller.AdminController
	var loginMiddleWare middleware.LoginRequired
	m.Group("/admin", func (r martini.Router)  {
		r.Get("", admin.Index)
		r.Get("/tech", admin.ArticleList);
	}, loginMiddleWare.Call)
}

//个人中心路由
func ucenterRoute(m *martini.ClassicMartini)  {
	var userCenter controller.UserCenterController
	m.Group("/ucenter",func (r martini.Router)  {
		//Get
		r.Get("/user/:id\\d{1,5}",userCenter.Index) 		//首页
		r.Get("/collect", userCenter.Collections)			//收藏管理
		r.Get("/sign_history", userCenter.SignLogHistory)	//签到记录
		r.Get("/points", userCenter.MyPoints)				//我的积分
		r.Get("/message",userCenter.Message)				//站内信息
		r.Get("/invite", userCenter.InviteFriends)			//邀请好友
		r.Get("/info", userCenter.SelfInfo)					//个人资料
		r.Get("/password", userCenter.ChangePassword)		//修改密码
		r.Get("/new_article", userCenter.NewArticleView)	//发表新文章view
		r.Get("/art_list/:status", userCenter.GetArtList)	//文章列表

		//POST
		r.Post("/api_check", userCenter.CheckDaily)			//每日签到
	})
}

//帖子路由
func articleRoute(m *martini.ClassicMartini)  {
	var articleController controller.ArticleController
	m.Group("/article",func(r martini.Router){
		m.Get("/:id\\d{1,5}", articleController.Detail)      		//文章详情
		m.Post("/api_new", articleController.NewArticle)				//发表文章
		m.Post("/api_prise", articleController.AddPriseNum)				//为文章点赞
		m.Post("/api_diss", articleController.AddDissNum)				//Diss一下文章
		m.Post("/api_delete", articleController.Delete)					//删除文章
		m.Post("/api_comment", articleController.AddReply)				//回复
		m.Post("/api_comment_delete", articleController.DeleteReply)    //删除回复
		m.Post("/comment_list", articleController.GetReplyList)			//评论列表
	})
}

//小分类路由
func subPathRoute(m *martini.ClassicMartini)  {
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
}
//404 not found
func route404(m *martini.ClassicMartini)  {
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
}