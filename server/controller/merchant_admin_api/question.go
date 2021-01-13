package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type Question struct {
}

func (c *Question) Get(ctx *gin.Context) {
}

func (c *Question) CreateQuestion(ctx *gin.Context) {
	var params struct {
		MerchantId     uint64
		QuestionTypeId uint64
		Question       string
		Answer         string
		Enable         bool
		Sort           uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.Questions{
		//Model:          models.Model{Id:&params.Id},
		MerchantId:     &params.MerchantId,
		QuestionTypeId: &params.QuestionTypeId,
		Question:       &params.Question,
		Answer:         &params.Answer,
		Enable:         &params.Enable,
		Sort:           &params.Sort,
	}
	res := global.DBMaster.Table(question.GetTableName()).Create(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
	return
}
func (c *Question) UpdateQuestion(ctx *gin.Context) {
	var params struct {
		Id             uint64
		MerchantId     uint64
		QuestionTypeId uint64
		Question       string
		Answer         string
		Enable         bool
		Sort           uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.Questions{
		Model:          models.Model{Id: &params.Id},
		MerchantId:     &params.MerchantId,
		QuestionTypeId: &params.QuestionTypeId,
		Question:       &params.Question,
		Answer:         &params.Answer,
		Enable:         &params.Enable,
		Sort:           &params.Sort,
	}
	res := global.DBMaster.Table(question.GetTableName()).Update(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
func (c *Question) DeleteQuestion(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.Questions{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBSlave.Table(question.GetTableName()).Delete(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "衫春成功")
	return

}
func (c *Question) CreateQuestionType(ctx *gin.Context) {
	var params struct {
		MerchantId       uint64
		QuestionTypeName string
		Enable           bool
		Sort             uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.QuestionTypes{
		//Model:            models.Model{Id: &params.Id},
		MerchantId:       &params.MerchantId,
		QuestionTypeName: &params.QuestionTypeName,
		Sort:             &params.Sort,
		Enable:           &params.Enable,
	}
	res := global.DBMaster.Table(question.GetTableName()).Create(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
	return
}
func (c *Question) UpdateQuestionType(ctx *gin.Context) {
	var params struct {
		Id               uint64
		MerchantId       uint64
		QuestionTypeName string
		Enable           bool
		Sort             uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.QuestionTypes{
		Model:            models.Model{Id: &params.Id},
		MerchantId:       &params.MerchantId,
		QuestionTypeName: &params.QuestionTypeName,
		Sort:             &params.Sort,
		Enable:           &params.Enable,
	}
	res := global.DBMaster.Table(question.GetTableName()).Update(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
func (c *Question) DeleteQuestionType(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	question := &models.QuestionTypes{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBSlave.Table(question.GetTableName()).Delete(question)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除承诺")
	return
}
