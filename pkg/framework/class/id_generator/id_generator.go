package id_generator

type IdGenerator interface {
	// 某个key进行累加
	Incr(key string) (int64, error)

	// 获取指定字段的值
	// 当目标数据不存在的时候，返回0，不报错
	Get(key string) (int64, error)

	// 哈希结构中，某个key进行累加
	HIncr(key string, field string) (int64, error)
	// 哈希结构中，获取某个key的当前值

	// 当目标数据不存在的时候，返回0，不报错
	HGet(key string, field string) (int64, error)

	// 哈希结构中，获取整个哈希结构的值
	HGetAll(key string) (map[string]int64, error)
}
