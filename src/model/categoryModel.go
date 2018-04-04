package model

import (
	"time"
)

//分类标签
type CategoryModel struct{
	Id int 
	Name string
	Pid int 
	Seq int 
	Url string
	Description string
	CreatedAt time.Time
}

func (CategoryModel) TableName() string {
	return "yz_category"
}