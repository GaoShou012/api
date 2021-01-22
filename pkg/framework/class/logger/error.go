package logger

type Error interface {
	// 错误时间
	Time() string

	// 错误ID
	Id() string

	// 错误的文件名
	Filename() string

	// 错误行号
	Line() int

	V() interface{}
	PushV(v interface{})

	// 输出错误信息
	Error() string
}
