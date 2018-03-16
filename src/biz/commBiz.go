package biz

import (
	"fmt"
	"model"
)

type CommomBiz struct{

}

func (this *CommomBiz) GetCategory(currentTab int) []model.CategoryModel {
	var category []model.CategoryModel
	err := GetDbInstance().Where("pid = ?", 0).Or("pid = ?", currentTab).Find(&category).Error
	if err != nil {
		fmt.Println(err)
	}
	return category
}

func GetHost() string{
	return "http://localhost:3000"
}