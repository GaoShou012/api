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

func (p *plugin) log(depth int, lv logger.Level, v ...interface{}) error {
	switch lv {
	case logger.ErrorLevel:
		glog.ErrorDepth(depth, v)
		break
	case logger.WarnLevel:
		glog.WarningDepth(depth, v)
		break
	case logger.InfoLevel:
		glog.InfoDepth(depth, v)
		break
	}
	return nil
}

func (p *plugin) Log(lv logger.Level, v ...interface{}) error {
	return p.log(2, lv, v)
}

func (p *plugin) Logf(lv logger.Level, format string, v ...interface{}) error {
	err := fmt.Errorf(format, v...)
	return p.log(2, lv, err)
}

func (p *plugin) Error(v ...interface{}) {
	p.log(2, logger.ErrorLevel, v...)
}

func (p *plugin) Warn(v ...interface{}) {
	p.log(2, logger.WarnLevel, v...)
}

func (p *plugin) Info(v ...interface{}) {
	p.log(2, logger.InfoLevel, v...)
}
