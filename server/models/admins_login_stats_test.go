package models

import (
	"api/config"
	"api/initialize"
	"testing"
)

func TestAdminsLoginStats_IncrByUserId(t *testing.T) {
	config.LocalLoad()
	if err := initialize.InitMysqlMaster(config.GetConfig().MysqlMaster); err != nil {
		panic(err)
	}
	m := &AdminsLoginStats{}
	res := m.IncrByUserId(1)
	if res.Error != nil {
		panic(res.Error)
	}
}
