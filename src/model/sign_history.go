package model

import (
	"time"
)

//签到记录表
type SignHistoryModel struct{
	Id int
	UserId int
	CreatedAt time.Time
}

func (SignHistoryModel) TableName() string {
	return "yz_sign_history"
}