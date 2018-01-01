package biz

import (
	"fmt"
	"model"
)
//帖子管理中心
type ArtManager struct{

}

//创建文章
func (manage ArtManager)  Create(article model.Article) int{
	err := GetDbInstance().Create(&article).Error
	if err != nil {
		return article.Id
	}
	return 0
}

//更新文章
func (manager ArtManager) Update(article model.Article, value interface{}) bool{
	err := GetDbInstance().Model(&article).Updates(value).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
//删除文章
func (manager ArtManager) Delete(article model.Article)  bool{
	err := GetDbInstance().Model(&article).Update("status", "D").Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//文章列表
func (manager ArtManager) GetArtList(artType int, limit int, offset int, status string) []model.Article{
	var article []model.Article
	err := GetDbInstance().Where("type = ? AND status = ?", artType, status).Limit(limit).Offset(offset).Find(&article).Error
	if err != nil {
		fmt.Println(err)
	}
	return article
}
//文章详情
func (manager ArtManager) Detail(artId int) model.Article{
	var article model.Article
	err := GetDbInstance().Where("id = ?", artId).First(&article).Error
	if err != nil {
		fmt.Println(err)
	}
	return article
}
