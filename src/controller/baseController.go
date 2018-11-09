package controller

import (
	"github.com/martini-contrib/sessions"
	"model"
	"biz"
	"common"
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

type commonResult struct {
	Code int
	Message string
	Result interface{}
}

var CommonResult commonResult
func GetUser(session sessions.Session) model.UserModel{
	v := session.Get("yz_session_token")
	var user model.UserModel
	sessionString, ok := v.(string)
	if ok && len(sessionString) > 0{
		user = biz.GetUserFromSession(sessionString)
	}
	return user
}

func Debug(v ...interface{}) {
	common.Debug(v)
}

func Error(v ...interface{}) {
	common.Error(v)
}

func SetCommonResult(code int, message string,result interface{}) {
	CommonResult.Code = code
	CommonResult.Message = message
	CommonResult.Result = result
}