package contracts

type Contract interface{}

type Contracts interface {
	// 创建合同
	Create(contract Contract) error

	// 审核合同
	Review(id uint64, args ...interface{}) error

	// 合同生效
	Effective(id uint64, args ...interface{}) error

	// 作废合同，无效合同
	Invalid(id uint64, args ...interface{}) error

	// 终止合同，相当于毁约
	Terminate(id uint64, args ...interface{}) error
}
