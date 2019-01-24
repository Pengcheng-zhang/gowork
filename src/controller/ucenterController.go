/*
 * @Author: pengcheng.zhang 
 * @Date: 2019-01-24 17:24:23 
 * @Last Modified by: pengcheng.zhang
 * @Last Modified time: 2019-01-24 17:29:08
 */
package controller

import (
	"biz"
	"fmt"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"model"
)

type UserCenterController struct {
}

type outPut struct {
	User model.UserModel
	ArtList []model.ArticleModel
	Categories interface{}
}

/**
 * @description: 首页
 * @methodType: GET
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) Index(r render.Render, session sessions.Session)  {
	var output outPut
	output.User = GetUser(session)
	r.HTML(200, "ucenter/index", output)
}

/**
 * @description: 下载记录
 * @methodType: GET
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) DownloadHistory(r render.Render, session sessions.Session)  {
	
}

/**
 * @description: 收藏管理
 * @methodType: GET
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) Collections(r render.Render, session sessions.Session)  {
	
}

/**
 * @description: 用户文章列表
 * @methodType: GET
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) GetArtList(r render.Render, session sessions.Session)  {
	var manager biz.ArtBiz
	var output outPut	
	output.User = GetUser(session)
	output.ArtList = manager.GetUserArtList(3, 10, 0, "A" ) //所有文章
	fmt.Printf("art list %v", output.ArtList)
	r.HTML(200, "ucenter/artlist", output)
}

/**
 * @description: 发表新贴 GET
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) NewArticleView(r render.Render, session sessions.Session)  {
	var categoryBiz biz.CategoryBiz
	var output outPut
	output.User = GetUser(session)
	output.Categories = categoryBiz.GetSubCategory()
	r.HTML(200, "article/new", output)
}

/**
 * @description: 用户设置
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) Settings(r render.Render,session sessions.Session) {
	var output outPut
	output.User = GetUser(session)
	r.HTML(200, "ucenter/settings", output)
}

/**
 * @description: 日常签到
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController)  ApiCheckDaily(r render.Render, session sessions.Session){
	var userBiz biz.UserBiz
	user := userBiz.GetCurrentUser()
	var checkmodel model.SignHistoryModel
	checkmodel.UserId = user.Id
	result := userBiz.CheckedIn(checkmodel)
	if result {
		setJsonResult(30001, "您今日已经签到过", "")
		r.JSON(200, jsonResult)
		return
	}
	result = userBiz.CheckIn(checkmodel)
	if result {
		setJsonResult(10000, "签到成功", "")
	}else{
		setJsonResult(30002, "签到失败", "")
	}
	r.JSON(200, jsonResult)
}


/**
 * @description: 我的积分
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) MyPoints(r render.Render, session sessions.Session)  {
	
}
/**
 * @description: 站内信息
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) Message(r render.Render, session sessions.Session)  {
	
}

/**
 * @description: 邀请好友
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) InviteFriends(r render.Render, session sessions.Session)  {
	
}

/**
 * @description: 个人资料
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) SelfInfo(r render.Render, session sessions.Session)  {
	
}

/**
 * @description: 签到记录
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) SignLogHistory(r render.Render, session sessions.Session)  {
	//user := userManager.GetCurrentUser()

}
/**
 * @description: 修改密码
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) ChangePassword(r render.Render, req *http.Request, session sessions.Session)  {
	var userBiz biz.UserBiz
	user := userBiz.GetCurrentUser()
	oldPassword := req.FormValue("old_password")
	newPassword := req.FormValue("new_password")
	//检查原始密码
	err := userBiz.CheckUserOldPassword(user.Username, oldPassword)
	if err != nil {
		setJsonResult(10003, "原始密码错误！", err)
		r.JSON(200, jsonResult)
		return
	}
	//更新新密码
	updateData := map[string]interface{}{"Password": newPassword}
	result,err := userBiz.UpdateUserInfo(user, updateData)
	if err != nil {
		setJsonResult(10002, "更改密码失败！", err)
	}else {
		setJsonResult(10000, "更新成功！", result)
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 更改用户名
 * @param {type} 
 * @return: 
 */
func (this *UserCenterController) ChangeUserName(r render.Render, req *http.Request, session sessions.Session)  {
	var userBiz biz.UserBiz
	user := userBiz.GetCurrentUser()
	newName := req.FormValue("username")
	if len(newName) > 20 || len(newName) == 0 {
		setJsonResult(10006, "用户名长度不符合！", "")
		r.JSON(200, jsonResult)
		return
	}
	//用户名特殊字符检查
	sensitive := userBiz.CheckSensitiveWord(newName)
	if sensitive == false {
		setJsonResult(10007, "用户名包含不合法字符！", "")
		r.JSON(200, jsonResult)
		return
	}
	//更新用户名
	updateData := map[string]interface{}{"username": newName}
	result,err := userBiz.UpdateUserInfo(user, updateData)
	if err != nil {
		setJsonResult(10005, "用户名更新失败！", err)
		r.JSON(200, jsonResult)
		return
	} else {
		setJsonResult(10000, "用户名更新成功！", result)
	}
	r.JSON(200, jsonResult)
}
