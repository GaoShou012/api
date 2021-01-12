package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	models_merchant "api/models/merchant"
	"github.com/gin-gonic/gin"
)

type Department struct {
}

//部门列表
func (c *Department) GetDepartmentList(ctx *gin.Accounts) {

}

//创建部门
func (c *Department) CreateDepartment(ctx *gin.Context) {
	var params struct {
		MerchantId     uint64
		DepartmentName string
		Sort           uint
		Desc           string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := &models_merchant.MerchantsDepartments{
		MerchantId:     &params.MerchantId,
		DepartmentName: &params.DepartmentName,
		Sort:           &params.Sort,
		Desc:           &params.Desc,
	}
	res := global.DBMaster.Table(department.GetTableName()).Create(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "创建成功")
}
func (c *Department) UpdateDepartment(ctx *gin.Context) {
	var params struct {
		Id             uint64
		MerchantId     uint64
		DepartmentName string
		Sort           uint
		Desc           string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := models_merchant.MerchantsDepartments{
		Model:          models.Model{Id: &params.Id},
		MerchantId:     &params.MerchantId,
		DepartmentName: &params.DepartmentName,
		Sort:           &params.Sort,
		Desc:           &params.Desc,
	}
	res := global.DBMaster.Table(department.GetTableName()).Update(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}
func (c *Department) DelDepartment(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := models_merchant.MerchantsDepartments{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBMaster.Table(department.GetTableName()).Delete(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
}

//部门下级小组
func (c *Department) CreateDepartmentGroup(ctx *gin.Context) {
	var params struct {
		MerchantId          uint64
		DepartmentGroupName string
		Sort                uint
		Desc                string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := &models_merchant.MerchantsDepartmentGroup{
		MerchantId:          &params.MerchantId,
		DepartmentGroupName: &params.DepartmentGroupName,
		Sort:                &params.Sort,
		Desc:                &params.Desc,
	}
	res := global.DBMaster.Table(department.GetTableName()).Create(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "创建成功")
}
func (c *Department) UpdateDepartmentGroup(ctx *gin.Context) {
	var params struct {
		Id                  uint64
		MerchantId          uint64
		DepartmentGroupName string
		Sort                uint
		Desc                string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := models_merchant.MerchantsDepartmentGroup{
		Model:               models.Model{Id: &params.Id},
		MerchantId:          &params.MerchantId,
		DepartmentGroupName: &params.DepartmentGroupName,
		Sort:                &params.Sort,
		Desc:                &params.Desc,
	}
	res := global.DBMaster.Table(department.GetTableName()).Update(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}
func (c *Department) DelDepartmentGroup(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	department := models_merchant.MerchantsDepartmentGroup{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBMaster.Table(department.GetTableName()).Delete(department)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
}
