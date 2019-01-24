package route

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"controller"
	"middleware"
	"model"
)

var loginMiddleWare middleware.LoginRequired
var adminMiddleWare middleware.AdminRequired
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
	uploadRoute()
	friendsRoute()
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
	mClassic.Get("/forget", home.Forget)		//忘记密码

	mClassic.Post("/api_signup", home.ApiRegist)  		//注册
	mClassic.Post("/api_login", home.ApiLogin)		//登陆
	mClassic.Post("/api_logout", home.ApiLogout)			//登出
	mClassic.Post("/api_forget_password", home.ApiForgetPassword)
}

//后台管理路由
func adminRoute()  {
	var admin controller.AdminController
	mClassic.Group("/admin", func (r martini.Router)  {
		r.Get("", admin.Index)
	}, loginMiddleWare.Call, adminMiddleWare.Call)
	mClassic.Group("/admin", func (r martini.Router)  {
		r.Post("/user_list", admin.UserList)
		r.Post("/article_list", admin.ArticleList)
	}, loginMiddleWare.Call, adminMiddleWare.Call)
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
		r.Post("/api_check", userCenter.ApiCheckDaily)			//每日签到
		
	}, loginMiddleWare.Call)
}

//帖子路由
func articleRoute()  {
	var articleController controller.ArticleController
	mClassic.Group("/article",func(r martini.Router){
		r.Get("/:id", articleController.GetDetail)      		//文章详情
		r.Post("/api_comment_list", articleController.ApiGetReplyList)			//评论列表
		r.Post("/api_detail", articleController.ApiDetail)
	})

	mClassic.Group("/article", func(r martini.Router){
		r.Post("/api_new", articleController.ApiNewArticle)				//发表文章
		r.Post("/api_update", articleController.ApiUpdate)				//更新文章
		r.Post("/api_prise", articleController.ApiAddPriseNum)				//为文章点赞
		r.Post("/api_diss", articleController.ApiAddDissNum)				//Diss一下文章
		r.Post("/api_delete", articleController.ApiDelete)					//删除文章
		r.Post("/api_add_comment", articleController.ApiAddReply)				//回复
		r.Post("/api_comment_delete", articleController.ApiDeleteReply)    //删除回复
		r.Post("/api_type_list", articleController.ApiGetTypeList)    //类型列表
	}, loginMiddleWare.Call)

	mClassic.Group("/article",func(r martini.Router){
		r.Post("/api_roll_back", articleController.ApiRollback)    //帖子打回
	}, adminMiddleWare.Call)
}
//分类路由
func catePathRoute() {
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

func uploadRoute() {
	var upload controller.UploadController
	mClassic.Group("/upload", func(r martini.Router) {
		r.Post("/file", upload.File)
	}, loginMiddleWare.Call)
}

func friendsRoute() {
	var friends controller.FriendsController
	mClassic.Group("/friends", func(r martini.Router) {
		r.Post("/list", friends.List)
		r.Post("/details", friends.Detail)
		r.Post("/new", friends.New)
		r.Post("/update", friends.Update)
	})
}
//404 not found
func route404()  {
	mClassic.NotFound(func(r render.Render, req *http.Request ) {
		if req.Method == "POST" {
			r.JSON(404, map[string]interface{}{"Code": 404, "Message" : "未找到请求接口"})
		}else {
			type output struct {
				User model.UserModel
			}
			var data output
			r.HTML(404, "404", data)
		}
	})
}