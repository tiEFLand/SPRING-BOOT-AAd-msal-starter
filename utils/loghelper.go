package utils

import (
	"log"
	"strings"
)

func GetLogger() *Logger {
	level, _ := GetConfig().Get("log.level")
	if level == "" {
		level