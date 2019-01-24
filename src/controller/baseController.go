package controller

import (
	"github.com/martini-contrib/sessions"
	"model"
	"biz"
	"services"
	"strconv"
)

type htmlResult struct {
	User model.UserModel
	CurrentCate model.CategoryModel
	Category []model.CategoryModel
	Articles []model.ArticleResultModel
	ArticleCount int
	ArticlePageCount int
	ArticleCurrentPage int
}

type jsonCommonResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Result interface{} `json:"result"`
}

var jsonResult jsonCommonResult
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
	services.Debug(v)
}

func Error(v ...interface{}) {
	services.Error(v)
}

func getPage(v string) (page int) {
	if v == "" {
		page = 0
	} else {
		var err error
		page,err = strconv.Atoi(v)
		if err != nil {
			page = 0
		}else {
			if page < 0 {
				page = 0
			}
		}
	}
	return page
}

func setJsonResult(code int, message string,result interface{}) {
	jsonResult.Code = code
	jsonResult.Message = message
	jsonResult.Result = result
}