package biz

import (
	"fmt"
	"model"
)

type History struct{

}

//记录管理员操作历史
func (history History) AddOperationHistory(opHistory model.OperationHistory) bool{
	err := GetDbInstance().Create(&opHistory).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}