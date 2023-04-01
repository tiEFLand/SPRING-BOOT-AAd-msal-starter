package utils

import (
	"log"
	"strings"
)

func GetLogger() *Logger {
	level, _ := GetConfig().Get("log.level")
	if level == "" {
		level = "info"
	}
	logger := NewLogger()
	debugHandle, err := NewRotatingFileHandler("server.log")
	if err != nil {
		log.Fatalf("打开日志文件失败：%+v", err)
	}
	if strings.EqualFold(level, "debug") {
		logger.SetLevel(LevelDebug)
	} else if strings.EqualFold(level, "info") {
		logger.SetLevel(LevelInfo)
	} else if strings.EqualFold(level, "warn") {
		logger.SetLevel(LevelWarn)
	} else if strings.EqualFold(level, "error") {
		logger.SetLevel(LevelError)
	} else if strings.EqualFold(level, "fatal") {
		logger.SetLevel(LevelFatal)
	}
	debugHandle.SetLevel(LevelDebug)
	logger.AddHandler(debugHandle)
	return logger
}
