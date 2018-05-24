package biz

import (
	"model"
)

type History struct{

}

//记录管理员操作历史
func (history History) AddOperationHistory(opHistory model.OperationHistoryModel) bool{
	err := GetDbInstance().Create(&opHistory).Error
	if err != nil {
		Debug("add operate history failed:", err.Error())
		return false
	}
	return true
}