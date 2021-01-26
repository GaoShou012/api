package logger_log

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
	v        []interface{}
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
func (e *Error) PushV(v interface{}) {
	e.v = append(e.v, v)
}

func (e *Error) Error() string {
	return fmt.Sprint("错误ID:", e.id, e.v)
}
