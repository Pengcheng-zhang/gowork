package biz

import (
	"model"
)

type CommomBiz struct{

}

func (this *CommomBiz) GetCategory(currentTab int) []model.CategoryModel {
	var categories []model.CategoryModel
	err := GetDbInstance().Where("pid = ?", 0).Or("pid = ?", currentTab).Find(&categories).Error
	if err != nil {
		Debug("get category failed:", err.Error())
	}
	return categories
}

func(this *CommomBiz) GetSubcateIds(pid int) []int {
	var categories []model.CategoryModel
	var cateIds []int
	err := GetDbInstance().Where("pid = ?", pid).Or("id = ?", pid).Find(&categories).Error
	if err != nil {
		Debug("get subcateids failed", err.Error())
		return cateIds
	}
	Debug("categories:",categories )
	for _,value := range categories {
		cateIds = append(cateIds, value.Id)
	}
	return cateIds
}

func (this *CommomBiz) GetCategoryByName(name string) model.CategoryModel{
	var category model.CategoryModel
	err := GetDbInstance().Where("url = ?", name).First(&category).Error
	if err != nil {
		Debug("get category by name failed", name, err.Error())
	}
	return category
}

func (this *CommomBiz) GetSubCategory() []model.CategoryModel {
	var categories []model.CategoryModel
	err := GetDbInstance().Where("pid > ?", 0).Find(&categories).Error
	if err != nil {
		Debug("get subcategory failed:", err.Error())
	}
	return categories
}