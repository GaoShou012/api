package rbac

type Middleware interface {
	/*
		保存操作者上下文
	*/
	SetOperator(ctx interface{}, operator Operator)
	/*
		获取操作者上下文
	*/
	GetOperator(ctx interface{}) (Operator, error)
	/*
		加密操作者信息
	*/
	Encrypt(key []byte, operator Operator) (string, error)
	/*
		解密操作者信息
	*/
	Decrypt(key []byte, str string) (Operator, error)
}
