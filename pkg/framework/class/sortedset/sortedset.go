package sortedset

type Sortedset interface {
	// sortedset是否存在
	Exists(topic string) (bool, error)
	// 获取整个sortedset数量
	Len(topic string) int64
	// 查询指定页面
	Find(topic string, page int64, pageSize int64) ([]Item, error)

	// 设置一个item
	SetItem(topic string, key string, val float64) error
	// 获取item位置，正向排序，距离最大数字偏移位置
	GetOffset(topic string, key string) (int64, error)
	// 获取item位置，负向排序，距离最少数字偏移位置
	GetOffsetN(topic string, key string) (int64, error)
	// 从正向获取item
	GetItemFormPositive(topic string) ([]Item, error)
	// 从反向获取item
	GetItemFromNegative(topic string) ([]Item, error)
	// item是否存在
	ExistsItem(topic string, key string) (bool, error)
}
