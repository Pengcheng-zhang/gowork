package controller

import (
	"biz"
	"fmt"
	"net/http"
	//"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	// "biz"
	"model"
)

type UserCenterController struct {
	output outPut
	userBiz biz.UserBiz
	commBiz biz.CommomBiz
	jResult interface{} //api请求返回结果
}

type outPut struct {
	User model.UserModel
	ArtList []model.ArticleModel
	Categories []model.CategoryModel
	Js []string
	Css []string
}

//首页
func (this *UserCenterController) Index(r render.Render, session sessions.Session)  {
	// if v == nil {
	// 	r.Redirect("/login")
	// }
	// user = biz.GetUserFromSession(v.(string))
	this.output.User = GetUser(session)
	this.output.Js = []string{}
	this.output.Css = []string{}
	r.HTML(200, "ucenter/index", this.output)
}
//下载记录
func (this *UserCenterController) DownloadHistory(r render.Render, session sessions.Session)  {
	
}
//收藏管理
func (this *UserCenterController) Collections(r render.Render, session sessions.Session)  {
	
}
//用户文章列表
func (this *UserCenterController) GetArtList(r render.Render, session sessions.Session)  {
	// v := session.Get("sucai_session_token")
	// fmt.Println(v)
	var user model.UserModel
	// if v == nil {
	// 	r.Redirect("/login")
	// }
	// user = biz.GetUserFromSession(v.(string))
	fmt.Println(user)
	var manager biz.ArtBiz	
	this.output.User = user
	this.output.ArtList = manager.GetUserArtList(3, 10, 0, "A" ) //所有文章
	fmt.Printf("art list %v", this.output.ArtList)
	this.output.Js = []string{}
	this.output.Css = []string{}
	r.HTML(200, "ucenter/artlist", this.output)
}
//日常签到
func (this *UserCenterController)  CheckDaily(r render.Render, session sessions.Session){
	user := this.userBiz.GetCurrentUser()
	var checkmodel model.SignHistoryModel
	checkmodel.UserId = user.Id
	result := this.userBiz.CheckedIn(checkmodel)
	if result {
		this.jResult = map[string]interface{}{"code": 30001, "message":"您今日已经签到过", "result": ""}
		r.JSON(200, this.jResult)
		return
	}
	result = this.userBiz.CheckIn(checkmodel)
	if result {
		this.jResult = map[string]interface{}{"code": 10000, "message":"签到成功", "result": ""}
	}else{
		this.jResult = map[string]interface{}{"code": 30002, "message":"签到失败", "result": ""}
	}
	r.JSON(200, this.jResult)
}
//发表新贴 GET
func (this *UserCenterController) NewArticleView(r render.Render, session sessions.Session)  {
	this.output.User = GetUser(session)
	this.output.Categories = this.commBiz.GetSubCategory()
	this.output.Js = []string{}
	this.output.Css = []string{}
	r.HTML(200, "article/new", this.output)
}
func (this *UserCenterController) Settings(r render.Render,session sessions.Session) {
	this.output.User = GetUser(session)
	r.HTML(200, "ucenter/settings", this.output)
}
//我的积分
func (this *UserCenterController) MyPoints(r render.Render, session sessions.Session)  {
	
}
//站内信息
func (this *UserCenterController) Message(r render.Render, session sessions.Session)  {
	
}
//邀请好友
func (this *UserCenterController) InviteFriends(r render.Render, session sessions.Session)  {
	
}
//个人资料
func (this *UserCenterController) SelfInfo(r render.Render, session sessions.Session)  {
	
}
//签到记录
func (this *UserCenterController) SignLogHistory(r render.Render, session sessions.Session)  {
	//user := userManager.GetCurrentUser()

}
//修改密码
func (this *UserCenterController) ChangePassword(r render.Render, req *http.Request, session sessions.Session)  {
	user := this.userBiz.GetCurrentUser()
	oldPassword := req.FormValue("old_password")
	newPassword := req.FormValue("new_password")
	//检查原始密码
	err := this.userBiz.CheckUserOldPassword(user.Username, oldPassword)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10003, "message" : "原始密码错误！", "result": err})
		return
	}
	//更新新密码
	updateData := map[string]interface{}{"Password": newPassword}
	result,err := this.userBiz.UpdateUserInfo(user, updateData)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10002, "message" : "更改密码失败！", "result": err})
		return
	}
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "更新成功！", "result": result})
}
//更改用户名
func (this *UserCenterController) ChangeUserName(r render.Render, req *http.Request, session sessions.Session)  {
	user := this.userBiz.GetCurrentUser()
	newName := req.FormValue("username")
	if len(newName) > 20 || len(newName) == 0 {
		r.JSON(200, map[string]interface{}{"code": 10006, "message" : "用户名长度不符合！", "result": ""})
		return
	}
	//用户名特殊字符检查
	sensitive := this.userBiz.CheckSensitiveWord(newName)
	if sensitive == false {
		r.JSON(200, map[string]interface{}{"code": 10007, "message" : "用户名包含不合法字符！", "result": ""})
		return
	}
	//更新用户名
	updateData := map[string]interface{}{"username": newName}
	result,err := this.userBiz.UpdateUserInfo(user, updateData)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10005, "message" : "用户名更新失败！", "result": err})
		return
	}
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "用户名更新成功！", "result": result})
}
