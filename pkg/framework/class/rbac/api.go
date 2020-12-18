package rbac

type Api interface {
	GetMethod() string
	GetPath() string
}

type ApiAdapter interface {
	Init() error

	Create(api Api) error
	//Delete(tenantId uint64, api Api) error
	Update(apiId uint64, api Api) error
	SelectById(tenantId uint64, apiId uint64) (Api, error)
	//SelectByPath(tenantId uint64, apiPath string) (Api, error)
	SelectByPage(tenantId uint64, page uint64, pageSize uint64) ([]Api, error)
	Count(tenantId uint64) (uint64, error)
	VerifyIdWithOperator(id uint64, operator Operator) (bool, error)
}
