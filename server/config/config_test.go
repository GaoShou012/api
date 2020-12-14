package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	{
		c := GetConfig()
		c.Load("./resource/dev.ini")
		fmt.Println("base:", c.Base)
		fmt.Println("mysql master:", c.MysqlMaster)
	}
}
