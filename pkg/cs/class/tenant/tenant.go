package tenant

import "cs/meta"

type Tenant interface {
	GetCustomerServers(code string) ([]meta.Client, error)
	NumberOfOpeningSessions(code string) (uint64, error)
}
