package model

import (
	"time"
)

type Category struct{
	Id int 
	Name string
	Pid int 
	Seq int 
	Url string
	CreatedAt time.Time
}

func (Category) TableName() string {
	return "yz_category"
}