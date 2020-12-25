package middleware

type Operator interface {
	SetContextId(uuid string)
	GetContextId() string
}

type OperatorContext interface {
	/*
		读取上下文信息
	*/
	Get(args ...interface{}) (Operator, error)

	/*
		签署上下文
	*/
	SignedString(args ...interface{}) (string, error)

	/*
		解析上下文
	*/
	Parse(args ...interface{}) interface{}

	/*
		上下文过期处理
	*/
	Expiration(args ...interface{}) interface{}
}
