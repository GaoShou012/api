package models

type QuestionType struct {
	Model
	CategoryId     *uint64
	MerchantId     *uint64
	Name           *string
	bindingSetting *int
	DialogueGroup  *string
	TenantId       *uint64
}

func (m *QuestionType) GetTableName() string {
	return "question_type"
}
