package models_merchant

import "api/models"

type MerchantsDepartmentGroup struct {
	models.Model
	MerchantId          *uint64 //商户id
	DepartmentGroupName *string //部门名
	Sort                *uint   //排序
	Desc                *string //描述
}

func (m *MerchantsDepartmentGroup) GetTableName() string {
	return "merchant_department_group"
}
