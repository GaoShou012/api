package env

import (
	"framework/class/logger"
	"framework/plugin/logger/logger_log"
)

var Logger logger.Logger

func init() {
	Logger = logger_log.New()
}
