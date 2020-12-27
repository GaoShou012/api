package models

import "time"

// 租户合同

type TenantsContracts struct {
	Model

	// 状态需要梳理，合同的流程，再设计出状态标记
	State          *uint64    // 合同状态 0=作废，1=不同过审核，2=已经确认生效，3=已经顺利完成，4=中途毁约
	TenantId       *uint64    // 租户ID
	EffectiveDate  *time.Time // 合同生效日期
	ExpirationDate *time.Time // 合同截止日期
	ShouldPay      *float64   // 应付金额
	RealPay        *float64   // 实付金额
	PaymentRecord  *string    // 支付记录，租户自定义填写

	// 审核人员
	ReviewerId   *string
	ReviewerName *string
	ReviewerTime *time.Time
	ReviewerNote *string // 审核人添加的合同备注

	// 合同创建人
	CreatorId   *uint64
	CreatorName *string
	CreatorNote *string // 创建人添加的合同备注
}
