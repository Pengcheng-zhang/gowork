package common

import (
	"os"
	"log"
	"time"
	"strings"
)

func getLogFile() *os.File{
	currentDate := time.Now().Format("2006-01-02")
	logFilePath := strings.Join([]string{"./log/", currentDate, ".log"}, "")
	_, err := os.Stat(logFilePath)
	if err != nil {
		logFile, fileErr := os.Create(logFilePath)
		if fileErr != nil{
			log.Println("create log file fail:", fileErr)
			return nil
		}
		return logFile
	}else {
		logFile,err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return nil
		}
		return logFile
	}
	
}

func Debug(v ...interface{}) {
	logFile := getLogFile()
	log.SetOutput(logFile)
	log.SetPrefix("[debug]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(v)
}

func Error(v ...interface{}) {
	logFile := getLogFile()
	log.SetOutput(logFile)
	log.SetPrefix("[error]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(v);
}