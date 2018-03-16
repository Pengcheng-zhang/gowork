package controller

import (
	"biz"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/sessions"
)

type AdminController struct{
	artBiz biz.ArtBiz
}

//后台管理首页
func (this *AdminController) Index(r render.Render, req *http.Request, session sessions.Session) {
	
}

//个人帖子列表
func (this *AdminController) ArticleList(r render.Render, req *http.Request, session sessions.Session) {
	
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