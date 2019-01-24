package services

import (
	"os"
	"log"
	"time"
	"strings"
	"fmt"
)

type LoggerService struct {
	
}
var g_logFile *os.File
func (this *LoggerService) initLoggerService() error{
	currentDate := time.Now().Format("2006-01-02")
	logFilePath := strings.Join([]string{"./log/", currentDate, ".log"}, "")
	_, err := os.Stat(logFilePath)
	if err != nil {
		g_logFile, err = os.Create(logFilePath)
		if err != nil{
			fmt.Println("create log file fail:", err)
			return err
		}
		fmt.Println("create log file name:", g_logFile.Name())
		return nil
	}else {
		g_logFile,err = os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println("open log file name:", g_logFile.Name())
		return nil
	}
	
}

func (this *LoggerService) GetLogFile() *os.File{
	if g_logFile != nil {
		return g_logFile
	}
	err := this.initLoggerService()
	if err == nil {
		return g_logFile
	}
	fmt.Println("get log file fail:", err)
	return nil
}

func Debug(v ...interface{}) {
	service := &LoggerService{}
	logFile := service.GetLogFile()
	if logFile != nil {
		log.SetOutput(logFile)
		log.SetPrefix("[debug]")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(v)
	}
}

func Error(v ...interface{}) {
	service := &LoggerService{}
	logFile := service.GetLogFile()
	if logFile != nil {
		log.SetOutput(logFile)
		log.SetPrefix("[error]")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(v);
	}
}