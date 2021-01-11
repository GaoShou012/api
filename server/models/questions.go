package models

type Questions struct {
	Model
	MerchantId     *uint64 //商户Id
	QuestionTypeId *uint64 //问题类型
	Question       *string //问题
	Answer         *string //问题答案
	Enable         *bool   //是否禁用
	Sort           *uint   //排序
}

func (m *Questions) GetTableName() string {
	return "questions"
}
