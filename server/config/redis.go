package config

type Redis struct {
	/*
		连接
	*/
	DNS string

	/*
		连接端口
	*/
	Port uint64

	/*
		验证密钥
	*/
	Password string

	/*
		连接池最大连接数量
	*/
	PoolMax uint64

	/*
		连接池最少连接数量
	*/
	PoolMin uint64
}
