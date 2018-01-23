package model

import (
	"time"
)

//管理员操作记录表
type OperationHistoryModel struct {
	Id int
	UserId int  //管理员id
	ArticleId int //帖子id
	Type string //操作类型 B：打回 D：删除
	Comment string //原因
	CreatedAt time.Time
}

func (OperationHistoryModel) TableName() string {
	return "yz_operation_history"
}