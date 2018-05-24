package biz

import (
	"errors"
	"time"
	"model"
)
//帖子管理中心
type ArtBiz struct{

}

type ArtResult struct{

}
//创建文章
func (this *ArtBiz)  Create(article model.ArticleModel) int{
	err := GetDbInstance().Create(&article).Error
	if err != nil {
		Debug("article create error:%s", err.Error())
		return 0
	}
	Debug("article id is : %d", article.Id)
	return article.Id
}

//更新文章
func (this *ArtBiz) Update(article model.ArticleModel, value interface{}) bool{
	err := GetDbInstance().Model(&article).Updates(value).Error
	if err != nil {
		Debug("update article failed:", err.Error())
		return false
	}
	return true
}

//删除文章
func (this *ArtBiz) Delete(article model.ArticleModel)  bool{
	err := GetDbInstance().Model(&article).Update("status", "D").Error
	if err != nil {
		Debug("delete article failed:", err.Error())
		return false
	}
	return true
}
//总数
func (this *ArtBiz) GetArtCount(artType int) int {
	var articleModel model.ArticleModel
	var count int = 0
	GetDbInstance().Where("type = ?", artType).Find(&articleModel).Count(&count)
	return count
}
//文章列表
func (this *ArtBiz) GetArtList(artType int, limit int, offset int, status string) []model.ArticleResultModel{
	var articles []model.ArticleResultModel
	selectField := "yz_tech.id, yz_tech.type, yz_tech.title, yz_tech.creator_id, left(yz_tech.content, 150) as content, "+
	"yz_tech.prise_num,yz_tech.diss_num, yz_tech.reply_num, yz_tech.view_times, yz_tech.last_reply_user_id,yz_tech.last_reply_time," +
	"yz_tech.status, yz_tech.created_at, yz_user.username as creator_name"
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech.creator_id"
	err := GetDbInstance().Table("yz_tech").Select(selectField).Joins(joinConditions).Where("yz_tech.type = ? AND status = ?", artType, status).Limit(limit).Offset(offset).Scan(&articles).Error
	if err != nil {
		Debug("get artlist failed:", err.Error())
	}
	return articles
}
//获取大分类及其下小分类文章
func (this *ArtBiz) GetTabArtList(types []int, limit int, offset int, status string) []model.ArticleResultModel{
	var articles []model.ArticleResultModel
	selectField := "yz_tech.id, yz_tech.type, yz_tech.title, yz_tech.creator_id, left(yz_tech.content, 150) as content, "+
	"yz_tech.prise_num,yz_tech.diss_num, yz_tech.reply_num, yz_tech.view_times, yz_tech.last_reply_user_id,yz_tech.last_reply_time," +
	"yz_tech.status, yz_tech.created_at, yz_user.username as creator_name"
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech.creator_id"
	err := GetDbInstance().Table("yz_tech").Select(selectField).Joins(joinConditions).Where("yz_tech.type IN (?) AND status = ?", types, status).Limit(limit).Offset(offset).Scan(&articles).Error
	if err != nil {
		Debug("get tab article list failed:", err.Error())
	}
	return articles
}

//用户文章列表
func (this *ArtBiz) GetUserArtList(userId int, limit int, offset int, status string)  []model.ArticleModel{
	var article []model.ArticleModel
	var err error
	if status == "A" {
		err = GetDbInstance().Where("creator_id = ?", userId).Limit(limit).Offset(offset).Find(&article).Error
	}else{
		err = GetDbInstance().Where("creator_id = ? AND status = ?", userId, status).Limit(limit).Offset(offset).Find(&article).Error
	}
	if err != nil {
		Debug("get user article list failed:", err.Error())
	}
	return article
}
//文章详情
func (this *ArtBiz) Detail(artId int) (model.ArticleModel,error){
	var article model.ArticleModel
	err := GetDbInstance().Where("id = ?", artId).First(&article).Error
	return article, err
}
//文章详情用于输出
func (this *ArtBiz) DetailOutput(artId int) (model.ArticleResultModel,error){
	var article model.ArticleResultModel
	selectFields := "yz_tech.*, yz_user.username as creator_name"
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech.creator_id"
	err := GetDbInstance().Table("yz_tech").Select(selectFields).Joins(joinConditions).Where("yz_tech.id = ?", artId).Scan(&article).Error
	return article, err
}

//添加评论
func (this *ArtBiz) AddReply(reply model.ReplyModel) error{
	article,err := this.Detail(reply.TechId)
	if err != nil || article.Id == 0{
		Debug("add reply failed:", err.Error())
		return err
	}
	err = GetDbInstance().Create(&reply).Error
	if err == nil {
		updateData := map[string]interface{}{"ReplyNum": article.ReplyNum+1,"LastReplyUserId": reply.UserId, "LastReplyTime": time.Now().Format("2006-01-02 15:04:05")}
		result := this.Update(article, updateData)
		if !result {
			return errors.New("更新错误")
		}
	}
	
	return err
}

//删除评论
func (this *ArtBiz) DeleteReply(reply model.ReplyModel) error{
	err := GetDbInstance().Model(&reply).Where("tech_id = ?", reply.TechId).Update("status", "D").Error
	return err
}

//获取回复列表
func (this *ArtBiz) GetReplyList(artId int, limit int, offset int) ([]model.ReplyResultModel,error){
	var replyList []model.ReplyResultModel
	article,err := this.Detail(artId)
	if err != nil || article.Id == 0 {
		Debug("get reply article detail failed:", err.Error())
		return replyList, err
	}
	selectFields := "yz_tech_reply.*, yz_user.username"
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech_reply.user_id"
	err = GetDbInstance().Table("yz_tech_reply").Select(selectFields).Joins(joinConditions).Where("tech_id = ?", artId).Limit(limit).Offset(offset).Scan(&replyList).Error
	if err != nil {
		Debug("get reply list failed", err.Error())
	}
	return replyList,err
}

//获取回复总数
func (this *ArtBiz) GetReplyCount(artId int) int {
	var replyModel []model.ReplyModel
	var count int = 0
	GetDbInstance().Where("tech_id = ?", artId).Find(&replyModel).Count(&count)
	return count
}