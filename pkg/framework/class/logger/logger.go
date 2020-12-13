package logger

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Init() error
	Log(level Level, v ...interface{}) error
	Logf(level Level, format string, v ...interface{}) error
	Name() string
	Destroy()
}
