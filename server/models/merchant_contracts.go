package models

import (
	"api/global"
	"fmt"
	"time"
)

const (
	MerchantsContractStateCreated = iota
	MerchantsContractStateReviewed
)

// 租户合同
type MerchantsContracts struct {
	Model

	// 状态需要梳理，合同的流程，再设计出状态标记
	State          *uint64    // 合同状态0=作废，1=审核中，2=通过审核，3=审核不通过，4=已经确认生效，5=已经顺利完成，6=中途毁约，7=逾期
	MerchantId     *uint64    // 租户ID
	EffectiveDate  *time.Time // 合同生效日期
	ExpirationDate *time.Time // 合同截止日期

	// 合同创建人
	CreatorId   *uint64
	CreatorName *string
	CreatorNote *string // 创建人添加的合同备注

	ShouldPayment *float64 // 应付金额
	RealPayment   *float64 // 实付金额
	PaymentNote   *string  // 支付备注，租户自定义填写

	// 审核人员
	ReviewerId   *uint64
	ReviewerName *string
	ReviewerTime *time.Time
	ReviewerNote *string // 审核人添加的合同备注
}

func (m *MerchantsContracts) GetTableName() string {
	return "merchants_contracts"
}

// 创建合同
func (m *MerchantsContracts) Create(
	tenantId uint64,
	effectiveDate time.Time, expirationDate time.Time,
	shouldPayment float64,
	creatorId uint64, creatorName string, creatorNote string,
) error {
	state := uint64(MerchantsContractStateCreated)
	i := &MerchantsContracts{
		Model:          Model{},
		State:          &state,
		MerchantId:     &tenantId,
		EffectiveDate:  &effectiveDate,
		ExpirationDate: &expirationDate,
		ShouldPayment:  &shouldPayment,
		RealPayment:    nil,
		PaymentNote:    nil,
		CreatorId:      &creatorId,
		CreatorName:    &creatorName,
		CreatorNote:    &creatorNote,
		ReviewerId:     nil,
		ReviewerName:   nil,
		ReviewerTime:   nil,
		ReviewerNote:   nil,
	}
	res := global.DBMaster.Table(m.GetTableName()).Create(i)
	return res.Error
}

// 审核合同
func (m *MerchantsContracts) review(id uint64, state uint64, reviewerId uint64, reviewerName string, reviewerNote string) error {
	//state := uint64(MerchantsContractStateReviewed)
	reviewerTime := time.Now()
	i := &MerchantsContracts{
		Model:          Model{},
		State:          &state,
		MerchantId:     nil,
		EffectiveDate:  nil,
		ExpirationDate: nil,
		CreatorId:      nil,
		CreatorName:    nil,
		CreatorNote:    nil,
		ShouldPayment:  nil,
		RealPayment:    nil,
		PaymentNote:    nil,
		ReviewerId:     &reviewerId,
		ReviewerName:   &reviewerName,
		ReviewerTime:   &reviewerTime,
		ReviewerNote:   &reviewerNote,
	}
	res := global.DBMaster.Table(m.GetTableName()).Where("id=?", id).Updates(i)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("审核失败，可能目标数据不存在，或者内容没有发生变化")
	}
	return nil
}

func (m *MerchantsContracts) GenReviewer(state uint64, reviewerId uint64, reviewerName string, reviewerNote string) interface{} {
	var args []interface{}
	args = append(args, state, reviewerId, reviewerName, reviewerNote)
	return args
}

func (m *MerchantsContracts) Review(id uint64, args ...interface{}) error {
	state := args[0].(uint64)
	reviewerId := args[1].(uint64)
	reviewName := args[2].(string)
	reviewerNote := args[3].(string)
	return m.review(id, state, reviewerId, reviewName, reviewerNote)
}

// 支付
func (m *MerchantsContracts) Payment(id uint64, realPayment float64, paymentNote string) error {
	i := &MerchantsContracts{
		Model:          Model{},
		State:          nil,
		MerchantId:     nil,
		EffectiveDate:  nil,
		ExpirationDate: nil,
		CreatorId:      nil,
		CreatorName:    nil,
		CreatorNote:    nil,
		ShouldPayment:  nil,
		RealPayment:    &realPayment,
		PaymentNote:    &paymentNote,
		ReviewerId:     nil,
		ReviewerName:   nil,
		ReviewerTime:   nil,
		ReviewerNote:   nil,
	}
	res := global.DBMaster.Table(m.GetTableName()).Where("id=?", id).Updates(i)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("支付失败")
	}
	return nil
}

// 作废合同，无效合同
func (m *MerchantsContracts) Invalid(id uint64, args ...interface{}) error {
	return nil
}
