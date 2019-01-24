package controller

import (
	"biz"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/sessions"
	"model"
	"strconv"
)

type AdminController struct{
	artBiz biz.ArtBiz
	userBiz biz.UserBiz
}

//后台管理首页
func (this *AdminController) Index(r render.Render, req *http.Request, session sessions.Session) {
	r.HTML(200, "admin/index", "", render.HTMLOptions{Layout: ""})
}

//后台用户列表
func (this *AdminController) UserList(r render.Render, req *http.Request, session sessions.Session) {
	in_page := req.FormValue("start")
	in_length := req.FormValue("length")
	type userList struct {
		Data []model.UserModel `json:"data"`
		RecordsFiltered int `json:"recordsFiltered"`
		RecordsTotal int `json:"recordsTotal"`
	}
	var output userList
	var page int = 0
	var length int = 10
	var err error
	if in_page == "" {
		page = 0
	} else {
		page,err = strconv.Atoi(in_page)
		if err != nil {
			page = 0
		}else {
			if page > 1 {
				page = page - 1
			}else {
				page = 0
			}
		}
	}

	if in_length != "" {
		length,err = strconv.Atoi(in_length)
		if err != nil {
			length = 10
		}
	}
	offset := page * length
	output.Data, output.RecordsTotal = this.userBiz.GetUserList(length, offset, "all")
	output.RecordsFiltered = output.RecordsTotal
	r.Header().Set("Access-Control-Allow-Origin", "*");
	r.JSON(200, output)
}
//帖子列表
func (this *AdminController) ArticleList(r render.Render, req *http.Request, session sessions.Session) {
	in_start := req.FormValue("start")
	in_length := req.FormValue("length")
	in_status := req.FormValue("status")
	in_search := req.FormValue("keyword")
	type articleList struct {
		Data []model.ArticleResultModel `json:"data"`
		RecordsFiltered int `json:"recordsFiltered"`
		RecordsTotal int `json:"recordsTotal"`
	}
	if in_status != "P" && in_status != "B" && in_status != "C" && in_status != "All" {
		in_status = "All"
	}
	var output articleList
	var offset int = getPage(in_start)
	var length int = 10
	var err error

	if in_length != "" {
		length,err = strconv.Atoi(in_length)
		if err != nil {
			length = 10
		}
	}
	output.Data, output.RecordsTotal = this.artBiz.GetArtList(0, length, offset, in_status, in_search)
	output.RecordsFiltered = output.RecordsTotal
	r.Header().Set("Access-Control-Allow-Origin", "*");
	r.JSON(200, output)
}

//发表文章
func (this *AdminController) PushArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}
//删除文章
func (this *AdminController) DeleteArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}

//更新文章
func (this *AdminController) UpdateArticle(r render.Render, req *http.Request, session sessions.Session) {

}

//不合规范帖子打回
func (this *AdminController) RollbackArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}