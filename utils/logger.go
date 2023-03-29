
package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// var l *log.Logger
// var Log = NewLogger()

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Logger struct {
	logs     map[string]*log.Logger
	handlers []IHandler
	level    Level
}

func NewLogger() *Logger {
	logger := &Logger{
		level:    LevelDebug,
		handlers: make([]IHandler, 0),
	}

	h, _ := NewStreamHandler()
	logger.handlers = append(logger.handlers, h)

	return logger
}

func toFmts(v ...interface{}) string {
	var fmtstr []string
	for _ = range v {
		fmtstr = append(fmtstr, "%v")
	}
	return strings.Join(fmtstr, " ")
}

func (this *Logger) AddHandler(handler IHandler) {
	this.handlers = append(this.handlers, handler)
}

func (this *Logger) SetLevel(level Level) {
	// this.level = level
	this.handlers[0].SetLevel(level)
}

func (this *Logger) Debug(v ...interface{}) {
	this.log(LevelDebug, toFmts(v...), v...)
}

func (this *Logger) Debugf(fmts string, v ...interface{}) {
	this.log(LevelDebug, fmts, v...)