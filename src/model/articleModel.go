package model

import (
	"time"
)

//文章表
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
	LastReplyTime string //最后回复时间
	CreatedAt time.Time
}

type ArticleResultModel struct{
	Id int `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	CreatorId int `json:"creator_id"`
	CreatorName string `json:"creator_name"`
	Content string  `json:"content"`
	PriseNum int   `json:"prise_num"`//点赞数
	DissNum int   `json:"diss_num"` //鄙视数
	ReplyNum int  `json:"reply_num"` //评论数
	ViewTimes int  `json:"view_times"` //阅读数
	LastReplyUserId int  `json:"last_reply_user_id"`//最后回复人
	Status string    `json:"status"`//C: 创建，P：发布，B:打回，S:保存 D：删除
	LastReplyTime string `json:"last_reply_time"` //最后回复时间
	CreatedAt time.Time `json:"created_at"`
}

func (ArticleModel) TableName() string {
	return "yz_tech"
}