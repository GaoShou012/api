package logger_glog

import (
	"fmt"
	"framework/class/logger"
	"github.com/golang/glog"
)

var _ logger.Logger = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Log(lv logger.Level, v ...interface{}) error {
	switch lv {
	case logger.ErrorLevel:
		glog.Error(v)
		break
	case logger.WarnLevel:
		glog.Warning(v)
		break
	case logger.InfoLevel:
		glog.Info(v)
		break
	}
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
