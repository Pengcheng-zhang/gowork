package biz

import (
	"services"
	"github.com/jinzhu/gorm"
)

func Debug(v ...interface{}) {
	services.Debug(v)
}

func Error(v ...interface{}) {
	services.Error(v)
}

func GetDbInstance() *gorm.DB{
	return services.GetDbInstance();
}