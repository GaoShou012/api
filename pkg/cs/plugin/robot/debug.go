package robot

import env2 "framework/env"

var debug bool

func Debug(enable bool) {
	debug = enable
}
func IsDebug() bool {
	if debug == false {
		return false
	}
	if env2.Logger.IsDebug() == false {
		return false
	}
	return true
}
