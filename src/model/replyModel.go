package model

import (
	"time"
)

type ReplyModel struct {
	Id int `json:"id"`
	TechId int `json:"article_id"`
	UserId int `json:"-"`
	Content string `json:"content"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ReplyResultModel struct {
	Id int `json:"id"`
	TechId int `json:"article_id"`
	UserId int `json:"-"`
	Username string `json:"user_name"`
	Content string `json:"content"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (ReplyModel) TableName() string {
	return "yz_tech_reply"
}