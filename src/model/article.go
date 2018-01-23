package model

import (
	"time"
)

type ArticleModel struct{
	Id int
	Type string
	Title string
	CreatorId int
	Content string 
	PriseNum int //点赞数
	DissNum int   //鄙视数
	ReplyNum int  //评论数
	ViewTimes int  //阅读数
	LastReplyUserId int //最后回复人
	Status string    //C: 创建，P：发布，B:打回，S:保存 D：删除
	LastReplyTime time.Time //最后回复时间
	CreatedAt time.Time
}

func (ArticleModel) TableName() string {
	return "yz_tech"
}