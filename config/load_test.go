package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	{
		c := GetConfig()
		c.Load("./database.ini")
		c.Load("./dev.ini")

		fmt.Println(c.Base, c.Mysql)
	}
}
