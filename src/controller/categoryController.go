package controller

import (
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"strconv"
	"biz"
)

type CategoryController struct {
	uCenterBiz biz.UserBiz
	commBiz biz.CommomBiz
	artBiz biz.ArtBiz
	hResult htmlResult  //html数据
	jResult interface{} //api请求返回结果
}

//分类
func (this *CategoryController) CategoryPath(r render.Render, req *http.Request, session sessions.Session) {
	p := req.FormValue("p")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}
	category := this.commBiz.GetCategoryByName(req.URL.Path)
	if category.Id == 0 {
		r.Redirect("/404")
		return
	}
	this.hResult.User = GetUser(session)
	this.hResult.Category = this.commBiz.GetCategory(category.Id)
	var cateIds []int
	for _,value := range this.hResult.Category {
		if value.Id == category.Id || value.Pid == category.Id {
			cateIds = append(cateIds, value.Id)
		}
	}
	this.hResult.CurrentCate = category
	this.hResult.Articles = this.artBiz.GetTabArtList(cateIds, 50, page, "P")
	r.HTML(200, "index", this.hResult)
}

func (this *CategoryController) SubCatePath(r render.Render, req *http.Request, session sessions.Session) {
	p := req.FormValue("p")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}
	category := this.commBiz.GetCategoryByName(req.URL.Path)
	if category.Id == 0 {
		r.Redirect("/404")
		return
	}
	this.hResult.CurrentCate = category
	this.hResult.User = GetUser(session)
	this.hResult.ArticleCount = this.artBiz.GetArtCount(category.Id)
	this.hResult.Articles = this.artBiz.GetArtList(category.Id, 50, page, "P")
	r.HTML(200, "subcate/index", this.hResult)
}