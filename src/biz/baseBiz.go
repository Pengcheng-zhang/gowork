package biz

import (
	"common"
	"github.com/jinzhu/gorm"
)

func Debug(v ...interface{}) {
	common.Debug(v)
}

func Error(v ...interface{}) {
	common.Error(v)
}

func GetDbInstance() *gorm.DB{
	return common.GetDbInstance();
}