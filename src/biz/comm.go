package biz

import (
	"fmt"
	"model"
)

type Commom struct{

}

func (comm Commom) GetCategory(currentTab int) []model.CategoryModel {
	var category []model.CategoryModel
	err := GetDbInstance().Where("pid = ?", 0).Or("pid = ?", currentTab).Find(&category).Error
	if err != nil {
		fmt.Println(err)
	}
	return category
}