package env

import (
	"framework/class/logger"
	"framework/plugin/logger/logger_glog"
)

var Logger logger.Logger

func init() {
	Logger = logger_glog.New()
}
