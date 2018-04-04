package controller

import (
	"github.com/martini-contrib/sessions"
	"model"
	"biz"
)

type htmlResult struct {
	Js []string
	Css []string
	User model.UserModel
	CurrentCate model.CategoryModel
	Category []model.CategoryModel
	Articles []model.ArticleResultModel
	ArticleCount int
}

func GetUser(session sessions.Session) model.UserModel{
	v := session.Get("yz_session_token")
	var user model.UserModel
	sessionString, ok := v.(string)
	if ok && len(sessionString) > 0{
		user = biz.GetUserFromSession(sessionString)
	}
	return user
}