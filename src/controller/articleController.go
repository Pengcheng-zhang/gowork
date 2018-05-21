package controller

import (
	"strconv"
	"biz"
	"strings"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	// "biz"
	"model"
)
type artOutPut struct{
	Article model.ArticleResultModel
	ReplyList []model.ReplyResultModel
	PageCount int
	ReplyCount int
	User model.UserModel
}

type ArticleController struct{
	artBiz biz.ArtBiz
	userBiz biz.UserBiz
	artOutput artOutPut
}

/********** ERROR CODE FOR ARTICLE ******************
error code 2XXXX
20001 帖子不存在
20002 创建失败
20003 更新失败
20004 删除失败
20005 点赞失败
20006 Diss失败
20007 
20008
20009
20010 标题不能为空
20011 类型不能为空
20012 内容不能为空
20013 标题过长
20014 
20020 数据错误
**************************END***********************/

//详情GET
func (this *ArticleController) Detail(r render.Render, params martini.Params, session sessions.Session) {
	artId := params["id"]
	replyPage := params["p"]
	id,err := strconv.Atoi(artId)
	if err != nil {
		fmt.Println("article id:", err)
		r.Redirect("/404")
		return
	}
	var page int = 1
	if replyPage == "" {
		page = 1
	} else {
		page,err = strconv.Atoi(replyPage)
		if err != nil {
			page = 1
		}
	}
	//获取主题内容
	article, err := this.artBiz.DetailOutput(id)
	if err != nil {
		r.Redirect("/404")
		return
	}
	//update view times
	updateData := map[string]interface{}{"ViewTimes": article.ViewTimes+1}
	var articleModel model.ArticleModel
	articleModel.Id = article.Id
	this.artBiz.Update(articleModel, updateData)

	this.artOutput.Article = article

	//获取对应的回复列表
	var offset int = (page - 1) * 50
	replyCount := this.artBiz.GetReplyCount(id)
	if offset > replyCount {
		offset = (replyCount / 50) * 50 
	}
	replyList, perr := this.artBiz.GetReplyList(id, 50, offset)
	if perr != nil {
		r.Redirect("/404")
		return
	}
	this.artOutput.User = GetUser(session)
	this.artOutput.ReplyCount = replyCount
	this.artOutput.PageCount = replyCount / 50
	this.artOutput.ReplyList = replyList
	r.HTML(200, "article/view", this.artOutput)
	
}

//创建/发表帖子 POST
func (this *ArticleController)  NewArticle(r render.Render, req *http.Request, session sessions.Session) {
	articleNode := req.FormValue("node")
	title := req.FormValue("title")
	content := req.FormValue("content")

	if articleNode == "" {
		r.JSON(200, map[string]interface{}{"code": 20011, "message" : "请选择主题节点"})
		return
	}
	if title == "" {
		r.JSON(200, map[string]interface{}{"code": 20010, "message" : "标题不能为空"})
		return
	}
	if strings.Count(title,"") -1 > 100 {
		r.JSON(200, map[string]interface{}{"code": 20013, "message" : "标题过长"})
		return
	}
	if content == "" {
		r.JSON(200, map[string]interface{}{"code": 20012, "message" : "内容不能为空"})
		return
	}
	user := GetUser(session)
	var article model.ArticleModel
	article.Title = title
	article.Type = articleNode
	article.Content = content
	article.CreatorId = user.Id
	article.Status = "P"

	artId := this.artBiz.Create(article)
	if artId > 0 {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "发表成功", "result": artId})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20020, "message" : "未知错误"})
	}
}

//更新帖子
func (this *ArticleController)  Update(r render.Render, req *http.Request, session sessions.Session){
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "未知错误"})
}

//帖子打回
func (this *ArticleController)  Rollback(r render.Render, req *http.Request, session sessions.Session){
	articleID := req.FormValue("article_id")
	fmt.Printf("article rollback id: %s", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	article, err:= this.artBiz.Detail(artId)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	updateData := map[string]interface{}{"Status": "B"}
	result := this.artBiz.Update(article, updateData)
	if result {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "打回成功"})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20006, "message" : "打回失败"})
	}
}

//删除帖子
func (this *ArticleController) Delete(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	article, err:= this.artBiz.Detail(artId)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	result := this.artBiz.Delete(article);
	if result {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "删除成功"})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20004, "message" : "删除失败"})
	}
}

//为帖子点赞
func (this *ArticleController) AddPriseNum(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("id")
	fmt.Println("articleID:",articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	article, err:= this.artBiz.Detail(artId)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	updateData := map[string]interface{}{"PriseNum": article.PriseNum+1}
	result := this.artBiz.Update(article, updateData);
	if result {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "点赞成功", "result": article.PriseNum+1})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20003, "message" : "点赞失败"})
	}
	
}

//Diss一下帖子
func (this *ArticleController) AddDissNum(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("id")
	fmt.Printf("article diss id: %s", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	article, err:= this.artBiz.Detail(artId)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	updateData := map[string]interface{}{"DissNum": article.DissNum+1}
	result := this.artBiz.Update(article, updateData);
	if result {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "Diss成功", "result": article.DissNum+1})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20006, "message" : "Diss失败"})
	}
}

//添加评论
func (this *ArticleController) AddReply(r render.Render, req *http.Request, session sessions.Session) {
	user := this.userBiz.GetCurrentUser()
	articleID := req.FormValue("id")
	replyContent := req.FormValue("content")
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	if len(replyContent) == 0 {
		r.JSON(200, map[string]interface{}{"code": 20021, "message" : "请输入评论内容"})
		return
	}
	article, err:= this.artBiz.Detail(artId)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	var reply model.ReplyModel
	reply.TechId = article.Id
	reply.Status = "A"
	reply.UserId = user.Id
	reply.Content = replyContent
	err = this.artBiz.AddReply(reply)
	if err == nil {
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "留言成功", "result": reply})
	}else{
		r.JSON(200, map[string]interface{}{"code": 20026, "message" : "留言失败"})
	}
}

//删除留言
func (this *ArticleController) DeleteReply(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	replyId := req.FormValue("reply_id")
	fmt.Printf("article delete reply id: %s,%s", articleID, replyId)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}

	repId,repErr := strconv.Atoi(replyId)
	if repErr != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "留言不存在"})
		return
	}
	var replyModel model.ReplyModel
	replyModel.Id = repId
	replyModel.TechId = artId
	err = this.artBiz.DeleteReply(replyModel)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "留言删除失败"})
		return
	}
	r.JSON(200, map[string]interface{}{"code": 10000, "message" : "留言删除成功"})
}

//留言列表
func (this *ArticleController) GetReplyList(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	limit := req.FormValue("limit")
	offset := req.FormValue("offset")
	fmt.Printf("article diss id: %s", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在"})
		return
	}
	article, errArt:= this.artBiz.Detail(artId)
	if errArt != nil {
		r.JSON(200, map[string]interface{}{"code": 20001, "message" : "文章不存在", "result": article.Id})
		return
	}
	limitInt,errLimit := strconv.Atoi(limit)
	if errLimit != nil {
		limitInt = 10
	}
	offsetInt,errOffset := strconv.Atoi(offset)
	if errOffset != nil {
		offsetInt = 0;
	}
	replyList,errList := this.artBiz.GetReplyList(artId, limitInt, offsetInt)
	if errList != nil {
		r.JSON(200, map[string]interface{}{"code": 20028, "message" : "获取失败", "result": replyList})
	}else{
		r.JSON(200, map[string]interface{}{"code": 10000, "message" : "", "result": replyList})
	}
}


