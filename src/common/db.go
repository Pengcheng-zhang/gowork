package common

import (
	"strings"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var dbInstance *gorm.DB
func  DbInit() *gorm.DB {
	databaseParams := GetConfigSection("database")
	db_connect_string := strings.Join([]string{databaseParams["user"], ":", databaseParams["password"], "@tcp(", databaseParams["host"], ":", databaseParams["port"], ")/", databaseParams["name"], "?charset=", databaseParams["charset"], "&parseTime=", databaseParams["parsetime"]}, "")

	dbInstance, _ = gorm.Open(databaseParams["driver"], db_connect_string)
	//defer MY_DB.Close()
	dbInstance.SetLogger(log.New(getLogFile(), "\r\n", 0))
	return dbInstance
}

func GetDbInstance()  *gorm.DB{
	return dbInstance;
}
