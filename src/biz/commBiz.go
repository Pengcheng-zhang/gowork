package biz

import (
	"fmt"
	"model"
)

type CommomBiz struct{

}

func (this *CommomBiz) GetCategory(currentTab int) []model.CategoryModel {
	var categories []model.CategoryModel
	err := GetDbInstance().Where("pid = ?", 0).Or("pid = ?", currentTab).Find(&categories).Error
	if err != nil {
		fmt.Println(err)
	}
	return categories
}

func(this *CommomBiz) GetSubcateIds(pid int) []int {
	var categories []model.CategoryModel
	err := GetDbInstance().Where("pid = ?", pid).Or("id = ?", pid).Find(&categories).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("categories:",categories )
	var cateIds []int
	for _,value := range categories {
		cateIds = append(cateIds, value.Id)
	}
	return cateIds
}

func (this *CommomBiz) GetCategoryByName(name string) model.CategoryModel{
	var category model.CategoryModel
	err := GetDbInstance().Where("url = ?", name).First(&category).Error
	if err != nil {
		fmt.Println(err)
	}
	return category
}

func (this *CommomBiz) GetSubCategory() []model.CategoryModel {
	var categories []model.CategoryModel
	err := GetDbInstance().Where("pid > ?", 0).Find(&categories).Error
	if err != nil {
		fmt.Println(err)
	}
	return categories
}

func GetHost() string{
	return "http://localhost:3000"
}