package models_merchant

import "api/models"

type MerchantsDepartment struct {
	models.Model
	MerchantId     *uint64
	ParentId       *uint64
	DepartmentName *string
	Sort           *uint
	Desc           *string
}
