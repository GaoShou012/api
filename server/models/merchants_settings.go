package models

type MerchantsSettings struct {
	Model
	MerchantId *uint64

	// 访客端的连接
	VisitorUrl *string

	// 访客的Token密本，用于给商户生成访客的token，然后使用此token，访问客服系统，识别访客的信息
	VisitorTokenCipherKey *string

	// Api密钥
	ApiCipherKey *string

	// 会话排队的最大数量
	MaxNumberOfSessionQueue *uint64

	// 是否优先分配给上一次对话过的客服
	EnableSessionAssignToTheCloseCs *bool
}

func (m *MerchantsSettings) SelectById(fields string,id uint64) (bool,error){
	
}