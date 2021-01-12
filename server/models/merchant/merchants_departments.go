package models_merchant

import "api/models"

type MerchantsDepartments struct {
	models.Model
	MerchantId     *uint64 //商户id
	DepartmentName *string //部门名
	Sort           *uint   //排序
	Desc           *string //描述
}

func (m *MerchantsDepartments) GetTableName() string {
	return "merchants_departments"
}
