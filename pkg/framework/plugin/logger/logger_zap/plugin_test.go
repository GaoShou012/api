package logger_zap

import (
	"framework/class/logger"
	"testing"
)

func TestPlugin_Log(t *testing.T) {
	p := New()
	p.Log(logger.ErrorLevel, "123", "456")
	p.Logf(logger.ErrorLevel, "fuck %s,%d", "123", 444)
}
