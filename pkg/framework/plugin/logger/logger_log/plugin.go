package logger_log

import (
	"fmt"
	"framework/class/logger"
	"log"
)

var _ logger.Logger = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Log(lv logger.Level, v ...interface{}) error {
	log.Println(lv, v)
	return nil
}

func (p *plugin) Logf(lv logger.Level, format string, v ...interface{}) error {
	err := fmt.Errorf(format, v...)
	return p.Log(lv, err)
}

func (p *plugin) Error(v ...interface{}) {
	p.Log(logger.ErrorLevel, v...)
}

func (p *plugin) Warn(v ...interface{}) {
	p.Log(logger.WarnLevel, v...)
}

func (p *plugin) Info(v ...interface{}) {
	p.Log(logger.InfoLevel, v...)
}
