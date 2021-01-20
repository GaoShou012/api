package logger_glog

import (
	"fmt"
	"framework/class/logger"
	"time"
)

var _ logger.Error = &Error{}

type Error struct {
	t        time.Time
	id       string
	fileName string
	line     int
	v        interface{}
}

func (e *Error) Time() string {
	return e.t.Format(time.RFC3339)
}

func (e *Error) Id() string {
	return e.id
}

func (e *Error) Filename() string {
	return e.fileName
}

func (e *Error) Line() int {
	return e.line
}

func (e *Error) V() interface{} {
	return e.v
}

func (e *Error) Error() string {
	return fmt.Sprintf("系统内部错误ID:%s", e.id)
}
