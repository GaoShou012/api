package env

import (
	"framework/class/logger"
	"framework/env"
)

var Logger logger.Logger

func init() {
	Logger = env.Logger
}
