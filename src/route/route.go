package route

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"controller"
	"middleware"
	"model"
)

var loginMiddleWare middleware.LoginRequired
var mClassic *martini.ClassicMartini
func Run(m *martini.ClassicMartini)  {
	mClassic = m
	commRoute()
	adminRoute()
	ucenterRoute()
	articleRoute()
	catePathRoute()
	wechatRoute()
	emailRoute()
	route404()

	m.Run()
}

//通用接口路由
func commRoute()  {
	var home controller.HomeController
	mClassic.Get("/", home.Index) 					//首页
	mClassic.Get("/signin", home.GetLogin)			//登陆页
	mClassic.Get("/signup", home.GetRegist)		//注册页
	mClassic.Get("/about", home.About)             //关于我们

	mClassic.Post("/api/signup", home.Regist)  		//注册
	mClassic.Post("/api/login", home.Login)		//登陆
	mClassic.Post("/api/logout", home.Logout)			//登出
}

//后台管理路由
func adminRoute()  {
	var admin controller.AdminController
	mClassic.Group("/admin", func (r martini.Router)  {
		r.Get("", admin.Index)
		r.Get("/tech", admin.ArticleList);
	}, loginMiddleWare.Call)
}

//个人中心路由
func ucenterRoute()  {
	var userCenter controller.UserCenterController
	mClassic.Group("/ucenter",func (r martini.Router)  {
		//Get
		r.Get("/index",userCenter.Index)   //个人中心
		r.Get("/user/:id",userCenter.Index) 		//用户
		r.Get("/collect", userCenter.Collections)			//收藏管理
		r.Get("/sign_history", userCenter.SignLogHistory)	//签到记录
		r.Get("/points", userCenter.MyPoints)				//我的积分
		r.Get("/message",userCenter.Message)				//站内信息
		r.Get("/invite", userCenter.InviteFriends)			//邀请好友
		r.Get("/info", userCenter.SelfInfo)					//个人资料
		r.Get("/password", userCenter.ChangePassword)		//修改密码
		r.Get("/new_article", userCenter.NewArticleView)	//发表新文章view
		r.Get("/art_list/:status", userCenter.GetArtList)	//文章列表
		r.Get("/settings", userCenter.Settings)
		//POST
		r.Post("/api_check", userCenter.CheckDaily)			//每日签到
		
	})
}

//帖子路由
func articleRoute()  {
	var articleController controller.ArticleController
	mClassic.Group("/article",func(r martini.Router){
		r.Get("/:id", articleController.Detail)      		//文章详情
		r.Post("/comment_list", articleController.GetReplyList)			//评论列表
	})

	mClassic.Group("/article", func(r martini.Router){
		r.Post("/api_new", articleController.NewArticle)				//发表文章
		r.Post("/api_prise", articleController.AddPriseNum)				//为文章点赞
		r.Post("/api_diss", articleController.AddDissNum)				//Diss一下文章
		r.Post("/api_delete", articleController.Delete)					//删除文章
		r.Post("/api_comment", articleController.AddReply)				//回复
		r.Post("/api_comment_delete", articleController.DeleteReply)    //删除回复
	}, loginMiddleWare.Call)
}
//分类路由
func catePathRoute() {
	var category controller.CategoryController
	mClassic.Get("/tab/\\w+", category.CategoryPath)
	mClassic.Get("/go/\\w+", category.SubCatePath)

}
//微信路由
func wechatRoute() {
	var wechat controller.WechatController
	mClassic.Group("/wechat", func(r martini.Router) {
		r.Get("", wechat.Index)
		r.Get("/login", wechat.Login)
	})
}

func emailRoute() {
	var email controller.EmailController
	mClassic.Group("/email", func(r martini.Router) {
		r.Get("/verify", email.Verification)
		r.Post("/send", email.SendRegistVerification)
	})
}
//404 not found
func route404()  {
	mClassic.NotFound(func(r render.Render) {
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