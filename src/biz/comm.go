package biz

import (
	"fmt"
	"model"
)

type Commom struct{

}

func (comm Commom) GetCategory(currentTab int) []model.Category {
	var category []model.Category
	err := GetDbInstance().Where("pid = ?", 0).Or("pid = ?", currentTab).Find(&category).Error
	if err != nil {
		fmt.Println(err)
	}
	return category
}