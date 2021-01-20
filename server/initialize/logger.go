package initialize

import (
	"api/global"
	"framework/plugin/logger/logger_log"
)

func InitLogger() {
	logger := logger_log.New()
	global.Logger = logger
}
