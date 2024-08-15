package log_demo

import (
	"log"
	"os"
)

func PrintLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	// 配置日志输出位置
	logFile, err := os.OpenFile("local.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %+v", err)
		return
	}

	log.SetOutput(logFile)
	log.Println("This is a log message")
	// 设置日志前缀
	log.SetPrefix("[prefix]")
	log.Println("This is a log message")
}

// 功能和 PrintLog 一样
func PrintLog1() {
	// 配置日志输出位置
	logFile, err := os.OpenFile("local.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %+v", err)
		return
	}

	logger := log.New(logFile, "[prefix]", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("This is a log message")
}
