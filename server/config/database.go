package config

type Mysql struct {
	/*
		连接DNS
	*/
	DNS string

	/*
		连接池最大连接
	*/
	PoolMax int

	/*
		连接池最少连接，空闲时保持的最少连接
	*/
	PoolMin int

	/*
		是否开启日志模式
	*/
	LogMode bool

	/*
		连接最大生存时间
	*/
	ConnMaxLifeTime uint64
}
