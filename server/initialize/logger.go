package initialize

import (
	"api/global"
	"framework/env"
)

func InitLogger() {
	global.Logger = env.Logger
}
