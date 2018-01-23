package output

import (
	"time"
)

//文章列表输出显示字段
type ArtlistResult struct{
	Id int
	Type string
	Title string
	CreatorId int
	CreatorName string //创建者名称
	Content string
	PriseNum int
	DissNum int
	ReplyNum int  //评论数
	ViewTimes int
	LastReplyUserId int
	Status string    //C: 创建，P：发布，B:打回，S:保存 D：删除
	LastReplyTime time.Time
	CreatedAt time.Time
}

//评论列表显示字段

