package stream_redis_stream

import (
	"framework/class/logger"
	"log"
)

var _ logger.Logger = &iLogger{}

type iLogger struct {
}

func (i iLogger) Init() error {
	return nil
}

func (i iLogger) Log(level logger.Level, v ...interface{}) error {
	switch level {
	case logger.ErrorLevel:
		log.Println("error:", v)
		break
	case logger.WarnLevel:
		log.Println("warn:", v)
		break
	case logger.InfoLevel:
		log.Println("info:", v)
		break
	default:
		log.Println("unknown:", v)
		break
	}
	return nil
}

func (i iLogger) Logf(level logger.Level, format string, v ...interface{}) error {
	panic("implement me")
}

func (i iLogger) Name() string {
	panic("implement me")
}

func (i iLogger) Destroy() {
	panic("implement me")
}
