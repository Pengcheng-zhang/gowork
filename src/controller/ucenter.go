package controller

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"biz"
	"model"
)

type UserCenter struct {

}

//首页
func (ucenter UserCenter) Index(r render.Render, session sessions.Session)  {
	v := session.Get("sucai_session_token")
	fmt.Println(v)
	var user model.User
	if v == nil {
		r.Redirect("/login")
	}
	user = biz.GetUserFromSession(v.(string))
	fmt.Println(user)
	type output struct {
		User model.User
		Js []string
		Css []string
	}
	var data output
	data.User = user
	data.Js = []string{"/scripts/dialog-plus-min.js","/js/main/common.js"}
	data.Css = []string{"/css/main/basev6.css","/css/ppt/workppt-publicv2.css","/css/ucenter/user.css"}
	r.HTML(200, "ucenter/index", data)
}
//下载记录
func (ucenter UserCenter) DownloadHistory(r render.Render, session sessions.Session)  {
	
}
//收藏管理
func (ucenter UserCenter) Collections(r render.Render, session sessions.Session)  {
	
}
//会员充值
func (ucenter UserCenter) Recharge(r render.Render, session sessions.Session)  {
	
}
//我的积分
func (ucenter UserCenter) MyPoints(r render.Render, session sessions.Session)  {
	
}
//站内信息
func (ucenter UserCenter) Message(r render.Render, session sessions.Session)  {
	
}
//邀请好友
func (ucenter UserCenter) InviteFriends(r render.Render, session sessions.Session)  {
	
}
//个人资料
func (ucenter UserCenter) SelfInfo(r render.Render, session sessions.Session)  {
	
}
//头像设置
func (ucenter UserCenter) PhotoSetting(r render.Render, session sessions.Session)  {
	
}
//修改密码
func (ucenter UserCenter) ChangePassword(r render.Render, session sessions.Session)  {
	
}
