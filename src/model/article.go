package model

import (
	"time"
)

type Article struct{
	Id int
	Type int
	Title int
	CreatorId int
	Content string
	PriseNum int
	DissNum int
	ViewTimes int
	LastReplyUserId int
	status string    //C: 创建，P：发布，B:打回，S:保存 D：删除
	LastReplyTime time.Time
	CreatedAt time.Time
}

func (Article) TableName() string {
	return "yz_tech"
}