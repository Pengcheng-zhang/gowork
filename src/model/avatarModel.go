package model

import (
	"time"
)

//文章表
type AvatarModel struct{
	Id int
	Url string
	Status string    //A: 有效图像，C：无效图像
	CreatedAt time.Time
}