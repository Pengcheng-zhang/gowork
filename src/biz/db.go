package biz

import (
	"strings"
	"github.com/jinzhu/gorm"
	"config"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var dbInstance *gorm.DB
func  DbInit() *gorm.DB {
	// var db_connect bytes.Buffer
	// db_connect.WriteString(config.DB_USER)
	// db_connect.WriteString(":")
	// db_connect.WriteString(config.DB_PASSWORD)
	// db_connect.WriteString("@tcp(")
	// db_connect.WriteString(config.DB_HOST)
	// db_connect.WriteString(":")
	// db_connect.WriteString(config.DB_PORT)
	// db_connect.WriteString(")/")
	// db_connect.WriteString(config.DB_NAME)
	// db_connect.WriteString("?charset=")
	// db_connect.WriteString(config.DB_CHARSET)
	// db_connect.WriteString("&parseTime=")
	// db_connect.WriteString(config.DB_PAESETIME)

	db_connect_string := strings.Join([]string{config.DB_USER, ":", config.DB_PASSWORD, "@tcp(", config.DB_HOST, ":", config.DB_PORT, ")/", config.DB_NAME, "?charset=", config.DB_CHARSET, "&parseTime=", config.DB_PAESETIME}, "")

	dbInstance, _ = gorm.Open(config.DB_DRIVER, db_connect_string)
	//defer MY_DB.Close()
	dbInstance.SetLogger(log.New(os.Stdout, "\r\n", 0))
	return dbInstance
}

func GetDbInstance()  *gorm.DB{
	return dbInstance;
}
