/*
 * @Description: 文章
 * @Author: pengcheng.zhang
 * @Date: 2018-10-30 19:43:10
 * @LastEditTime: 2019-01-25 00:21:33
 * @LastEditors: Please set LastEditors
 */

package controller

import (
	"strconv"
	"biz"
	"strings"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"model"
)

type ArticleController struct{
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
/**
 * @description: 列表
 * @methodType: POST
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiGetList(r render.Render, req *http.Request) {
	in_page := req.FormValue("p")
	type articleList struct {
		Articles []model.ArticleResultModel `json:"articles"`
		ArticleCount int `json:"total_count"`
		ArticlePageCount int `json:"page_count"`
		ArticleCurrentPage int `json:"current_page"`
	}
	var output articleList
	var artBiz biz.ArtBiz
	var page int = getPage(in_page)
	
	offset := page * 50
	output.Articles, output.ArticleCount = artBiz.GetArtList(0, 50, offset, "P", "")
	output.ArticlePageCount = output.ArticleCount / 50 + 1
	output.ArticleCurrentPage = page
	setJsonResult(10000, "success", output)
	r.JSON(200, jsonResult)
}

/**
 * @description: 详情GET
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) GetDetail(r render.Render, params martini.Params, session sessions.Session) {
	artId := params["id"]
	replyPage := params["p"]
	id,err := strconv.Atoi(artId)
	if err != nil {
		Error("article id:", id, err.Error())
		r.Redirect("/404")
		return
	}
	var page int = getPage(replyPage)
	//获取主题内容
	var artBiz biz.ArtBiz
	article, err := artBiz.DetailOutput(id)
	if err != nil {
		Error("get article detail failed:", err.Error())
		r.Redirect("/404")
		return
	}
	type DetailOutput struct{
		Article model.ArticleResultModel
		ReplyList []model.ReplyResultModel
		PageCount int
		ReplyCount int
		User model.UserModel
	}
	var detailOutput DetailOutput
	//update view times
	updateData := map[string]interface{}{"ViewTimes": article.ViewTimes+1}
	var articleModel model.ArticleModel
	articleModel.Id = article.Id
	artBiz.Update(articleModel, updateData)

	detailOutput.Article = article

	//获取对应的回复列表
	var offset int = (page - 1) * 50
	replyCount := artBiz.GetReplyCount(id)
	if offset > replyCount {
		offset = (replyCount / 50) * 50 
	}
	replyList, perr := artBiz.GetReplyList(id, 50, offset)
	if perr != nil {
		Error("get article reply failed:", err.Error())
		r.Redirect("/404")
		return
	}
	detailOutput.User = GetUser(session)
	detailOutput.ReplyCount = replyCount
	detailOutput.PageCount = replyCount / 50
	detailOutput.ReplyList = replyList
	r.HTML(200, "article/view", detailOutput)
}

/**
 * @description: 文章详情
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiDetail(r render.Render, req *http.Request) {
	artId := req.FormValue("id")
	id,err := strconv.Atoi(artId)
	if err != nil {
		setJsonResult(20011, "未找到文章", "")
		r.JSON(200, jsonResult)
		return
	}
	//获取主题内容
	var artBiz biz.ArtBiz
	article, err := artBiz.DetailOutput(id)
	if err != nil {
		setJsonResult(20011, "未找到文章", "")
	} else {
		setJsonResult(10000, "success", article)
	}
	r.JSON(200, jsonResult)
}
/**
 * @description: 创建/发表帖子 POST
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController)  ApiNewArticle(r render.Render, req *http.Request, session sessions.Session) {
	articleType := req.FormValue("type")
	title := req.FormValue("title")
	content := req.FormValue("content")

	if articleType == "" {
		setJsonResult(20011, "请选择文章主题", "")
		r.JSON(200, jsonResult)
		return
	}
	if title == "" {
		setJsonResult(20010, "标题不能为空", "")
		r.JSON(200, jsonResult)
		return
	}
	if strings.Count(title,"") -1 > 100 {
		setJsonResult(20013, "标题过长", "")
		r.JSON(200, jsonResult)
		return
	}
	if content == "" {
		setJsonResult(20012, "内容不能为空", "")
		r.JSON(200, jsonResult)
		return
	}
	artType, err := strconv.Atoi(articleType)
	if err != nil {
		setJsonResult(20014, "文章类型错误，重新选择", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	err = artBiz.CheckArticleType(artType)
	if err != nil {
		setJsonResult(20014, "文章类型错误，重新选择", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	var article model.ArticleModel
	article.Title = title
	article.Type = articleType
	article.Content = content
	article.CreatorId = user.Id
	article.Status = "P"

	artId := artBiz.Create(article)
	if artId > 0 {
		setJsonResult(10000, "发表成功", artId)
	}else{
		setJsonResult(20020, "未知错误", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 更新帖子
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController)  ApiUpdate(r render.Render, req *http.Request, session sessions.Session){
	articleId := req.FormValue("id")
	articleType := req.FormValue("type")
	title := req.FormValue("title")
	content := req.FormValue("content")

	if articleType == "" {
		setJsonResult(20011, "请选择文章主题", "")
		r.JSON(200, jsonResult)
		return
	}
	if title == "" {
		setJsonResult(20010, "标题不能为空", "")
		r.JSON(200, jsonResult)
		return
	}
	if strings.Count(title,"") -1 > 100 {
		setJsonResult(20013, "标题过长", "")
		r.JSON(200, jsonResult)
		return
	}
	if content == "" {
		setJsonResult(20012, "内容不能为空", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	id,_ := strconv.Atoi(articleId)
	article, err := artBiz.Detail(id)
	if err != nil {
		setJsonResult(20012, "未找到该文章，不能更新", "")
		r.JSON(200, jsonResult)
		return
	}
	artType, err := strconv.Atoi(articleType)
	if err != nil {
		setJsonResult(20014, "文章类型错误，重新选择", "")
		r.JSON(200, jsonResult)
		return
	}
	err = artBiz.CheckArticleType(artType)
	if err != nil {
		setJsonResult(20014, "文章类型错误，重新选择", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	if user.Id != article.CreatorId {
		setJsonResult(20013, "你不是文章作者，不能更新", "")
		r.JSON(200, jsonResult)
		return
	}

	updateData := map[string]interface{}{"title": title, "type": articleType, "content": content}
	err = artBiz.Update(article, updateData)
	if err == nil {
		setJsonResult(10000, "更新成功", id)
	}else{
		setJsonResult(20020, "未知错误", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 帖子打回
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController)  ApiRollback(r render.Render, req *http.Request, session sessions.Session){
	articleID := req.FormValue("article_id")
	Debug("article rollback id: ", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	article, err:= artBiz.Detail(artId)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	updateData := map[string]interface{}{"Status": "B"}
	err = artBiz.Update(article, updateData)
	if err == nil {
		setJsonResult(10000, "打回成功", "")
	}else{
		setJsonResult(20006, "打回失败", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 删除帖子
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiDelete(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	var artBiz biz.ArtBiz
	article, err:= artBiz.Detail(artId)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	if user.Id != article.CreatorId || strings.Index(user.Roles, "A") != -1{
		setJsonResult(20004, "你不是文章作者，无权删除", "")
		r.JSON(200, jsonResult)
		return
	}
	err = artBiz.Delete(article);
	if err == nil{
		setJsonResult(10000, "删除成功", "")
	}else{
		setJsonResult(20004, "删除失败", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 为帖子点赞
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiAddPriseNum(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("id")
	Debug("articleID:",articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	article, err:= artBiz.Detail(artId)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	result := artBiz.CheckDiss(article.Id, user.Id)
	if result.Id > 0 {
		message := "兄弟，你已经赞过了！"
		if result.ClickType == "D" {
			message = "兄弟，你已经Diss过了！"
		}
		setJsonResult(20004, message, "")
		r.JSON(200, jsonResult)
		return
	}
	result.ArticleId = article.Id
	result.UserId = user.Id
	result.ClickType = "P"
	err = artBiz.AddDissHistory(result)
	if err != nil {
		setJsonResult(20003, "点赞失败", "")
	}else {
		updateData := map[string]interface{}{"PriseNum": article.PriseNum+1}
		err = artBiz.Update(article, updateData);
		if err == nil {
			setJsonResult(10000, "点赞成功", article.PriseNum+1)
		}else{
			setJsonResult(20003, "点赞失败", "")
		}
	}
	
	r.JSON(200, jsonResult)
}

/**
 * @description: Diss一下帖子
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiAddDissNum(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("id")
	Debug("article diss id: ", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	article, err:= artBiz.Detail(artId)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	result := artBiz.CheckDiss(article.Id, user.Id)
	if result.Id > 0 {
		message := "兄弟，你已经Diss过了！"
		if result.ClickType == "P" {
			message = "兄弟，你已经赞过了！"
		}
		setJsonResult(20004, message, "")
		r.JSON(200, jsonResult)
		return
	}
	result.ArticleId = article.Id
	result.UserId = user.Id
	result.ClickType = "D"
	err = artBiz.AddDissHistory(result)
	if err != nil {
		setJsonResult(20003, "Diss失败", "")
	}else {
		updateData := map[string]interface{}{"DissNum": article.DissNum+1}
		err = artBiz.Update(article, updateData);
		if err == nil {
			setJsonResult(10000, "Diss成功", article.DissNum+1)
		}else{
			setJsonResult(20006, "Diss失败", "")
		}
	}
	
	r.JSON(200, jsonResult)
}

/**
 * @description: 添加评论
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiAddReply(r render.Render, req *http.Request, session sessions.Session) {
	user := GetUser(session)
	articleID := req.FormValue("id")
	replyContent := req.FormValue("content")
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	if len(replyContent) == 0 {
		setJsonResult(20001, "请输入评论内容", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	article, err:= artBiz.Detail(artId)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var reply model.ReplyModel
	reply.TechId = article.Id
	reply.Status = "A"
	reply.UserId = user.Id
	reply.Content = replyContent
	err = artBiz.AddReply(reply)
	if err == nil {
		setJsonResult(10000, "留言成功", reply)
	}else{
		setJsonResult(20026, "留言失败", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 删除留言
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiDeleteReply(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	replyId := req.FormValue("reply_id")
	Debug("article delete reply id:", articleID, replyId)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}

	repId,repErr := strconv.Atoi(replyId)
	if repErr != nil {
		setJsonResult(20001, "留言不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var replyModel model.ReplyModel
	replyModel.Id = repId
	replyModel.TechId = artId
	var artBiz biz.ArtBiz
	replyModel, err = artBiz.GetReply(replyModel)
	if err != nil {
		setJsonResult(20001, "留言删除失败", "")
		r.JSON(200, jsonResult)
		return
	}
	user := GetUser(session)
	if user.Id != replyModel.UserId || strings.Index(user.Roles, "A") != -1{
		setJsonResult(20004, "该留言不是你所写，无权删除", "")
		r.JSON(200, jsonResult)
		return
	}
	err = artBiz.DeleteReply(replyModel)
	if err != nil {
		setJsonResult(20001, "留言删除失败", "")
	} else {
		setJsonResult(10000, "留言删除成功", "")
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 留言列表
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiGetReplyList(r render.Render, req *http.Request, session sessions.Session) {
	articleID := req.FormValue("article_id")
	limit := req.FormValue("limit")
	offset := req.FormValue("offset")
	Debug("article diss id:", articleID)
	artId,err := strconv.Atoi(articleID)
	if err != nil {
		setJsonResult(20001, "文章不存在", "")
		r.JSON(200, jsonResult)
		return
	}
	var artBiz biz.ArtBiz
	article, errArt:= artBiz.Detail(artId)
	if errArt != nil {
		setJsonResult(20001, "文章不存在", article.Id)
		r.JSON(200, jsonResult)
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
	replyList,errList := artBiz.GetReplyList(artId, limitInt, offsetInt)
	if errList != nil {
		setJsonResult(20028, "获取失败", "")
	}else{
		setJsonResult(10000, "success", replyList)
	}
	r.JSON(200, jsonResult)
}

/**
 * @description: 获取文章类型列表
 * @methodType: 
 * @param {type} 
 * @return: 
 */
func (this *ArticleController) ApiGetTypeList(r render.Render) {
	var artBiz biz.ArtBiz
	result, err := artBiz.GetTypeList()
	if err != nil {
		setJsonResult(20029, "获取文章类型失败", "")
	} else {
		setJsonResult(10000, "success", result)
	}
	r.JSON(200, jsonResult)
}


