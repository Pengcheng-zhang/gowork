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
	r render.Render
	session sessions.Session 
	req *http.Request
}

type outPut struct {
	User model.UserModel
	ArtList []model.ArticleModel
	Js []string
	Css []string
}

var output outPut
var userManager biz.UCenterManager

//首页
func (this *UserCenterController) Index()  {
	// v := session.Get("sucai_session_token")
	// fmt.Println(v)
	var user model.UserModel
	// if v == nil {
	// 	r.Redirect("/login")
	// }
	// user = biz.GetUserFromSession(v.(string))
	fmt.Println(user)

	output.User = user
	output.Js = []string{}
	output.Css = []string{}
	this.r.HTML(200, "ucenter/index", output)
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
	var manager biz.ArtManager	
	output.User = user
	output.ArtList = manager.GetUserArtList(3, 10, 0, "A" ) //所有文章
	fmt.Printf("art list %v", output.ArtList)
	output.Js = []string{}
	output.Css = []string{}
	r.HTML(200, "ucenter/artlist", output)
}
//日常签到
func (this *UserCenterController)  CheckDaily(r render.Render, session sessions.Session){
	user := userManager.GetCurrentUser()
	var checkmodel model.SignHistoryModel
	checkmodel.UserId = user.Id
	result := userManager.CheckedIn(checkmodel)
	if result {
		jResult = map[string]interface{}{"code": 30001, "message":"您今日已经签到过", "result": ""}
		r.JSON(200, jResult)
		return
	}
	result = userManager.CheckIn(checkmodel)
	if result {
		jResult = map[string]interface{}{"code": 10000, "message":"签到成功", "result": ""}
	}else{
		jResult = map[string]interface{}{"code": 30002, "message":"签到失败", "result": ""}
	}
	r.JSON(200, jResult)
}
//发表新贴 GET
func (this *UserCenterController) NewArticleView(r render.Render)  {
	var user model.UserModel
	fmt.Println(user)
	output.User = user
	output.Js = []string{}
	output.Css = []string{}
	r.HTML(200, "ucenter/new_article", output)
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
	user := userManager.GetCurrentUser()
	oldPassword := req.FormValue("old_password")
	newPassword := req.FormValue("new_password")
	//检查原始密码
	err := userManager.CheckUserOldPassword(user.Username, oldPassword)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10003, "message" : "原始密码错误！", "result": err})
		return
	}
	//更新新密码
	updateData := map[string]interface{}{"Password": newPassword}
	result,err := userManager.UpdateUserInfo(user, updateData)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10002, "message" : "更改密码失败！", "result": err})
		return
	}
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "更新成功！", "result": result})
}
//更改用户名
func (this *UserCenterController) ChangeUserName(r render.Render, req *http.Request, session sessions.Session)  {
	user := userManager.GetCurrentUser()
	newName := req.FormValue("username")
	if len(newName) > 20 || len(newName) == 0 {
		r.JSON(200, map[string]interface{}{"code": 10006, "message" : "用户名长度不符合！", "result": ""})
		return
	}
	//用户名特殊字符检查
	sensitive := userManager.CheckSensitiveWord(newName)
	if sensitive == false {
		r.JSON(200, map[string]interface{}{"code": 10007, "message" : "用户名包含不合法字符！", "result": ""})
		return
	}
	//更新用户名
	updateData := map[string]interface{}{"username": newName}
	result,err := userManager.UpdateUserInfo(user, updateData)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 10005, "message" : "用户名更新失败！", "result": err})
		return
	}
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "用户名更新成功！", "result": result})
}
