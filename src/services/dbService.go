package services

import (
	"strings"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
)

type DbService struct{
}
var GlobalDbInstance *gorm.DB

func initDbService(){
	databaseParams := GetConfigSection("database")
	db_connect_string := strings.Join([]string{databaseParams["user"], ":", databaseParams["password"], "@tcp(", databaseParams["host"], ":", databaseParams["port"], ")/", databaseParams["name"], "?charset=", databaseParams["charset"], "&parseTime=", databaseParams["parsetime"]}, "")
	fmt.Println("db connect")
	GlobalDbInstance, _ = gorm.Open(databaseParams["driver"], db_connect_string)
	// defer this.DbInstance.Close()
	logger := &LoggerService{}
	GlobalDbInstance.SetLogger(log.New(logger.GetLogFile(), "\r\n", 0))
	GlobalDbInstance.LogMode(true)
}

func GetDbInstance()  *gorm.DB{
	if GlobalDbInstance == nil {
		initDbService()
	}
	return GlobalDbInstance
}
