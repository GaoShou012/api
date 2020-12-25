package rbac

type Model interface {
	/*
		获取Table名字
	*/
	GetTableName() string

	/*
		更新前执行
		可以清理ID，UpdatedAt，CreatedAt 等等
	*/
	BeforeUpdate()
}
