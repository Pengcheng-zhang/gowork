/*
 * @Description: 文章点赞、鄙视历史记录
 * @Author: pengcheng.zhang
 * @Date: 2019-01-24 23:52:30
 * @LastEditTime: 2019-01-25 00:13:59
 * @LastEditors: Please set LastEditors
 */
package model

import (
	"time"
)

//文章表
type PriseHistoryModel struct{
	Id int `json:"id"`
	ArticleId int `json:"article_id"`
	UserId	int `json:"user_id"`
	ClickType string   `json:"click_type"` //P: prise，D：diss
	CreatedAt time.Time `json:"created_at"`
}

func (PriseHistoryModel) TableName() string {
	return "yz_prise_history"
}