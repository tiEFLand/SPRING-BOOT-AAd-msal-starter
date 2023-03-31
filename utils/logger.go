
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
}

func (this *Logger) Info(v ...interface{}) {
	var fmtstr []string
	for _ = range v {
		fmtstr = append(fmtstr, "%v")
	}
	this.log(LevelInfo, strings.Join(fmtstr, " "), v...)
}

func (this *Logger) Infof(fmts string, v ...interface{}) {
	this.log(LevelInfo, fmts, v...)
}

func (this *Logger) Warn(v ...interface{}) {
	this.log(LevelWarn, toFmts(v...), v...)
}

func (this *Logger) Warnf(fmts string, v ...interface{}) {
	this.log(LevelWarn, fmts, v...)
}

func (this *Logger) Error(v ...interface{}) {
	this.log(LevelError, toFmts(v...), v...)
}

func (this *Logger) Errorf(fmts string, v ...interface{}) {
	this.log(LevelError, fmts, v...)
}

func (this *Logger) Fatal(v ...interface{}) {
	this.log(LevelFatal, toFmts(v...), v...)
}

func (this *Logger) Fatalf(fmts string, v ...interface{}) {
	this.log(LevelFatal, fmts, v...)
}

func (this *Logger) log(level Level, fmts string, v ...interface{}) {
	_, filename, line, ok := runtime.Caller(2)
	content := fmt.Sprintf(fmts, v...)
	if ok {
		filenames := strings.Split(filename, "/")
		filename = filenames[len(filenames)-1]
	}
	levelFmt := ""
	switch level {
	case LevelDebug:
		{
			levelFmt = "[DBUG]"
		}
	case LevelInfo:
		{
			levelFmt = "[INFO]"
		}
	case LevelWarn:
		{
			levelFmt = "[WARN]"
		}
	case LevelError:
		{
			levelFmt = "[ERRO]"
		}
	case LevelFatal:
		{
			levelFmt = "[FTAL]"
		}
	}

	for _, h := range this.handlers {
		if h.getLevel() > level {
			continue
		}
		if level == LevelFatal {
			h.logger().Fatalf("[%s:%d\t] %s %s", filename, line, levelFmt, content)
		} else {
			h.logger().Printf("[%s:%d\t] %s %s", filename, line, levelFmt, content)
		}

	}
}

type IHandler interface {
	SetLevel(Level)
	getLevel() Level
	logger() *log.Logger
}

type RotatingFileHandler struct {
	log     *log.Logger
	level   Level
	logPath string
}

func NewRotatingFileHandler(logPath string) (IHandler, error) {
	h := &RotatingFileHandler{
		level:   LevelDebug,
		logPath: logPath,
	}

	file, err := os.OpenFile(
		logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644,
	)
	if err != nil {
		return nil, err
	}

	h.log = log.New(file, "", log.LstdFlags)

	var i IHandler
	i = h
	return i, nil
}

func (this *RotatingFileHandler) SetLevel(level Level) {
	this.level = level
}

func (this *RotatingFileHandler) getLevel() Level {
	return this.level
}

func (this *RotatingFileHandler) logger() *log.Logger {
	return this.log
}

type StreamHandler struct {
	log   *log.Logger
	level Level
}

func NewStreamHandler() (IHandler, error) {
	h := &StreamHandler{
		level: LevelDebug,
	}

	h.log = log.New(os.Stdout, "", log.LstdFlags)

	var i IHandler
	i = h
	return i, nil
}

func (this *StreamHandler) SetLevel(level Level) {
	this.level = level
}

func (this *StreamHandler) getLevel() Level {
	return this.level
}

func (this *StreamHandler) logger() *log.Logger {
	return this.log
}