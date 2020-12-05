package logs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

/*
	日志等级设置
	日志等级: 0 关闭，1使命，2警告，3信息，4调试
	日志接口: Fatal | Warning | Info
*/
var logLevel = 4

func init() {
	SetLogLevel(4)
	//设置log格式为json
	log.SetFormatter(&log.JSONFormatter{})
	//log设置为可写入文件的模式
	log.SetOutput(os.Stdout)
	//设置等级
	log.SetLevel(SelectLogLevel(logLevel))
}

//选择报错级别默认为调试模式
func SelectLogLevel(setLevel int) (level log.Level) {
	switch setLevel {
	case 4:
		return log.DebugLevel
	case 3:
		return log.InfoLevel
	case 2:
		return log.WarnLevel
	case 1:
		return log.FatalLevel
	default:
		return log.DebugLevel
	}
}

//取日志等级
func GetLogLevel() int {
	return logLevel
}

//设置日志等级
func SetLogLevel(lev int) {
	logLevel = lev
}
func Fatal(err error, message string) {
	log.WithFields(log.Fields{
		"msg":  message,
		"error": err,
	}).Fatal("发生致命错误Th")

}
func Warning(err error, message string) {
	log.WithFields(log.Fields{
		"msg":    message,
		"error": err,
	}).Warn("警告错误")
}
func Info(err error, message string) {
	log.WithFields(log.Fields{
		"msg": message,
		"error":   err,
	}).Info("消息提醒")
}
