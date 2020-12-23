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
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
}
