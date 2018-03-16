package biz

import (
	"fmt"
	"model"
	"output"
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
		fmt.Printf("article create error:%s", err)
		return 0
	}
	fmt.Printf("article id is : %d", article.Id)
	return article.Id
}

//更新文章
func (this *ArtBiz) Update(article model.ArticleModel, value interface{}) bool{
	err := GetDbInstance().Model(&article).Updates(value).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//删除文章
func (this *ArtBiz) Delete(article model.ArticleModel)  bool{
	err := GetDbInstance().Model(&article).Update("status", "D").Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//文章列表
func (this *ArtBiz) GetArtList(artType int, limit int, offset int, status string) []output.ArtlistResult{
	var article []output.ArtlistResult
	selectField := `"yz_tech.id, yz_tech.type, yz_tech.title, yz_tech.creator_id, left(yz_tech.content, 150) as content, 
	yz_tech.prise_num,yz_tech.diss_num, yz_tech.view_times, yz_tech.last_reply_user_id,yz_tech.last_reply_time,
	yz_tech.status, yz_tech.created_at, yz_user.username as creator_name"`
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech.creator_id"
	err := GetDbInstance().Table("yz_tech").Select(selectField).Joins(joinConditions).Where("type = ? AND status = ?", artType, status).Limit(limit).Offset(offset).Scan(&article).Error
	if err != nil {
		fmt.Println(err)
	}
	return article
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
		fmt.Println(err)
	}
	return article
}
//文章详情
func (this *ArtBiz) Detail(artId int) (model.ArticleModel,error){
	var article model.ArticleModel
	err := GetDbInstance().Where("id = ?", artId).First(&article).Error
	fmt.Printf("article detail id:%d", article.Id)
	return article, err
}

//添加评论
func (this *ArtBiz) AddReply(reply model.ReplyModel) error{
	article,err := this.Detail(reply.TechId)
	if err != nil || article.Id == 0{
		fmt.Println(err)
		return err
	}
	err = GetDbInstance().Create(&reply).Error
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
		fmt.Println(err)
		return replyList, err
	}
	selectFields := "yz_tech_reply.*, yz_user.username"
	joinConditions := "JOIN yz_user ON yz_user.id = yz_tech_reply.user_id"
	err = GetDbInstance().Table("yz_tech_reply").Select(selectFields).Joins(joinConditions).Where("tech_id = ?", artId).Limit(limit).Offset(offset).Scan(&replyList).Error
	if err != nil {
		fmt.Println(err)
	}
	return replyList,err
}

//获取回复总数
func (this *ArtBiz) GetReplyCount(artId int) int {
	var replyModel model.ReplyModel
	var count int = 0
	GetDbInstance().Where("tech_id = ?", artId).Find(&replyModel).Count(&count)
	return count
}