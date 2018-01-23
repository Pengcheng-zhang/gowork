package model

import (
	"time"
)

type CategoryModel struct{
	Id int 
	Name string
	Pid int 
	Seq int 
	Url string
	CreatedAt time.Time
}

func (CategoryModel) TableName() string {
	return "yz_category"
}