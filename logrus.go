package llogger

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	logruslog "github.com/sirupsen/logrus"
	"github.com/natefinch/lumberjack"
)

var (
	LLogger *logruslog.Logger
)

var ErrNilLogger = errors.New("Logger is nil,please check the whether logger is init correct! ")

type LogConf struct {
	// FileName 日志文件路径
	FileName string `json:"fileName" yaml:"fileName"`
	// MaxSize 每个日志文件保存的最大尺寸 单位:M
	MaxSize int `json:"maxSize" yaml:"maxSize"`
	// MaxBackups 日志文件最多保存多少个备份
	MaxBackups int `json:"maxBackups" yaml:"maxBackups"`
	// MaxAge 文件最多保存多少天
	MaxAge int `json:"maxAge"  yaml:"maxAge"`
	// Compress 是否压缩
	Compress bool `json:"compress" yaml:"compress"`
	// Console 是否打印到控制台
	Console bool `json:"console" yaml:"console"`
}

// 初始化全局的Logrus日志框架
func InitLogger(c *LogConf) error {
	// TODO 传入配置的校验
	if c == nil {
		c = &LogConf{
			FileName:   "./log/default.log",
			MaxSize:    100,
			MaxBackups: 1,
			MaxAge:     30,
			Compress:   false}
	}
	LLogger = logruslog.New()

	if !c.Console {
		LLogger.SetFormatter(&MyFormatter{PrettyPrint: false})
		SetWriter(c)
	} else {
		LLogger.SetFormatter(&MyFormatter{PrettyPrint: true})
	}
	return nil
}

func SetWriter(c *LogConf) {
	lumberLog := &lumberjack.Logger{
		Filename:   c.FileName,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}
	LLogger.SetOutput(lumberLog)
}


func GetLogger() *logruslog.Logger {
	if LLogger == nil {
		log.Fatal("init logrus error")
	}

	return LLogger
}
func Info(message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprint(message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
	}).Info(msg)
}

func Infof(format string, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprintf(format, message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
	}).Info(msg)
}

func Warn(err error, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprint(message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Warn(err)
}

func Warnf(err error, format string, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprintf(format, message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Warn(err)
}

// Fatal日志基本打印时Logger会自动退出
func Fatal(err error, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprint(message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Fatal(err)
}

// Fatalf日志基本打印时Logger会自动退出
func Fatalf(err error, format string, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprintf(format, message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Fatal(err)
}

func Error(err error, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprint(message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Error(err)
}

func Errorf(err error, format string, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprintf(format, message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
		"msg": msg,
	}).Error(err)
}

func Debug(message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprint(message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
	}).Debug(msg)
}

func Debugf(format string, message ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	msg := fmt.Sprintf(format, message...)
	GetLogger().WithFields(logruslog.Fields{
		"files":   files,
	}).Debug(msg)
}