package models

type MerchantsSettings struct {
	Model
	MerchantId *uint64

	// 访客端的连接
	CustomerUrl *string
	// Api密钥
	ApiCipherKey *string

	// 会话排队的最大数量
	MaxNumberOfSessionQueue *uint64

	// 是否优先分配给上一次对话过的客服
	EnableSessionAssignToTheCloseCs *bool
}
