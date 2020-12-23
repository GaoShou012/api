package rbac

type Api interface {
	GetId() uint64
	GetTableName() string
	GetMethod() string
	GetPath() string
}

type ApiAdapter interface {
	Init() error

	/*
		校验操作者，是否有权限操作API
	*/
	Authority(operator Operator, apiId uint64) (bool, error)

	Create(api Api) error
	Delete(apiId uint64) error
	Update(apiId uint64, api Api) error
	SelectById(apiId uint64) (Api, error)

	/*
		根据ID，查询API
	*/
	FindById(operator Operator, apiId uint64, api Api) error
	FindByPage(operator Operator, page uint64, pageSize uint64, res []Api) error
	Count(tenantId uint64) (uint64, error)
}
