package logger_zap

import (
	"framework/class/logger"
	"go.uber.org/zap"
)

var _ logger.Logger = &plugin{}

type plugin struct {
	*zap.SugaredLogger
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Log(lv logger.Level, v ...interface{}) error {
	switch lv {
	case logger.ErrorLevel:
		p.SugaredLogger.Error(v...)
		break
	case logger.WarnLevel:
		p.SugaredLogger.Warn(v...)
		break
	case logger.InfoLevel:
		p.SugaredLogger.Info(v...)
		break
	default:
		p.SugaredLogger.Fatal(v...)
		break
	}
	return nil
}

func (p *plugin) Logf(lv logger.Level, format string, v ...interface{}) error {
	switch lv {
	case logger.ErrorLevel:
		p.SugaredLogger.Errorf(format, v...)
		break
	case logger.WarnLevel:
		p.SugaredLogger.Warnf(format, v...)
	case logger.InfoLevel:
		p.SugaredLogger.Infof(format, v...)
	default:
		p.SugaredLogger.Fatalf(format, v...)
		break
	}
	return nil
}

func (p *plugin) Error(v ...interface{}) {
	p.Log(logger.ErrorLevel, v)
}

func (p *plugin) Warn(v ...interface{}) {
	p.Log(logger.WarnLevel, v)
}

func (p *plugin) Info(v ...interface{}) {
	p.Log(logger.InfoLevel, v)
}
