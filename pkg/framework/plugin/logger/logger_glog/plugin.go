package logger_glog

import (
	"fmt"
	"framework/class/logger"
	"github.com/golang/glog"
	uuid "github.com/satori/go.uuid"
	"runtime"
	"time"
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

func (p *plugin) Error(v ...interface{}) logger.Error {
	err, ok := v[0].(logger.Error)
	if ok {
		return err
	}
	_, filename, line, _ := runtime.Caller(1)
	err = &Error{
		t:        time.Now(),
		id:       uuid.NewV4().String(),
		fileName: filename,
		line:     line,
		v:        v,
	}
	arr := make([]interface{}, len(v)+1)
	arr[0] = fmt.Sprintf("错误ID:%s", err.Id())
	for key, val := range v {
		arr[key+1] = val
	}
	fmt.Printf("%s:%s,Id:%s\n%s,%d\nerr=%v\n", logger.ErrorLevel, err.Time(), err.Id(), err.Filename(), err.Line(), err.V())
	return err
}

func (p *plugin) Warn(v ...interface{}) {
	p.log(2, logger.WarnLevel, v...)
}

func (p *plugin) Info(v ...interface{}) {
	p.log(2, logger.InfoLevel, v...)
}
