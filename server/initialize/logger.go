package initialize

import (
	"api/global"
	"framework/plugin/logger/logger_glog"
)

func InitLogger() {
	logger := logger_glog.New()
	global.Logger = logger
}
