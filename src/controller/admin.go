package controller

import (
	"biz"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/sessions"
)

type Admin struct{

}

var artManager biz.ArtManager
//后台管理首页
func (admin Admin) Index(r render.Render, req *http.Request, session sessions.Session) {
	
}

//帖子列表
func (admin Admin) ArticleList(r render.Render, req *http.Request, session sessions.Session) {
	
}

//发表文章
func (admin Admin) PushArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}
//删除文章
func (admin Admin) DeleteArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}

//更新文章
func (admin Admin) UpdateArticle(r render.Render, req *http.Request, session sessions.Session) {

}

//不合规范帖子打回
func (admin Admin) RollbackArticle(r render.Render, req *http.Request, session sessions.Session)  {
	
}