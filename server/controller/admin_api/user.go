package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/libs/validator"
	"api/models"
	"github.com/gin-gonic/gin"
)

type User struct{}

//创建用户
func (c *User) Insert(ctx *gin.Context) {
	var params struct {
		UserType uint64
		Username string
	}

	userType := params.UserType
	username := params.Username
	password := "123456"

	// validator
	{

		if err := libs_validator.Username(username); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}

		if err := libs_validator.Password(password); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	// save
	{
		admin := models.Admins{
			Id:        nil,
			Enable:    nil,
			State:     nil,
			UserType:  &userType,
			Username:  &username,
			Password:  &password,
			Nickname:  nil,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
		res := global.DBMaster.Create(admin)
		if res.Error != nil {
			libs_http.RspState(ctx, 1, res.Error)
			return
		}
	}

	libs_http.RspState(ctx, 0, "创建成功")
}

//删除用户
func (c *User) Delete() {

}

//修改用户
func (c *User) Update() {

}

//用户信息  列表
func (c *User) Select() {

}
