/*
 * @Description: 文章类型model
 * @Author: pengcheng.zhang
 * @Date: 2019-01-24 18:53:23
 * @LastEditTime: 2019-01-24 18:55:45
 * @LastEditors: Please set LastEditors
 */
package model

import (
	"time"
)

//文章表
type ArticleTypeModel struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Status string   `json:"status"` //A: 有效图像，C：无效图像
	CreatedAt time.Time `json:"-"`
}

func (ArticleTypeModel) TableName() string {
	return "yz_tech_type"
}