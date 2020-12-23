package rbac

type Operator interface {
	GetTenantId() uint64
	//GetId() uint64
	//GetUsername() string
}
