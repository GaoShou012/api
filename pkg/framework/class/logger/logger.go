package logger

type Level int8

func (l Level) String() string {
	switch l {
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	default:
		return "Unknown"
	}
}

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
)

type Logger interface {
	Init() error
	Log(level Level, v ...interface{}) error
	Logf(level Level, format string, v ...interface{}) error
	Error(v ...interface{}) Error

	// 隐藏错误输出到日志，并返回错误ID信息
	//ErrorInternal(v ...interface{}) Error

	Warn(v ...interface{})
	Info(v ...interface{})
}
