package models

type QuestionTypes struct {
	Model
	MerchantId       *uint64 //商户id
	QuestionTypeName *string //问题类型名称
	Sort             *uint   //排序
	Enable           *bool   //是否禁用
}

func (m *QuestionTypes) GetTableName() string {
	return "question_types"
}
