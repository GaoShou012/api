package models

import (
	"api/config"
	"api/global"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestAdmins(t *testing.T) {
	config.LocalLoad()
	if err := global.InitMysqlMaster(config.GetConfig().MysqlMaster); err != nil {
		log.Error(err)
	}

	userType := uint64(1)
	username := "gaoshou"
	password := "123"
	nickname := "123123"

	admin := &Admins{
		Id:        nil,
		Enable:    nil,
		State:     nil,
		UserType:  &userType,
		Username:  &username,
		Password:  &password,
		Nickname:  &nickname,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	// 添加
	{
		res := global.DBMaster.Create(admin)
		if res.Error != nil {
			panic(res.Error)
		}
	}

	// 更新
	{
		enable := true
		//state := uint64(1)
		admin := &Admins{
			Id:        nil,
			Enable:    &enable,
			State:     nil,
			UserType:  nil,
			Username:  nil,
			Password:  nil,
			Nickname:  nil,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
		res := global.DBMaster.Model(admin).Where("id=?", 1).Updates(admin)
		if res.Error != nil {
			panic(res.Error)
		}
	}

	// 查询
	{
		admin := &Admins{}
		res := global.DBMaster.First(admin, "id=?", 1)
		if res.Error != nil {
			panic(res.Error)
		}
		fmt.Println(admin)
	}

	{
		admin := &Admins{}
		res := global.DBMaster.Delete(admin, "id=?", 2)
		if res.Error != nil {
			panic(res.Error)
		}
	}
}
