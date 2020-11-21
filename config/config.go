package config

import (
	"gopkg.in/ini.v1"
	"sync"
)

var configInit sync.Once
var config *Config

func GetConfig() *Config {
	configInit.Do(func() {
		config = &Config{}
	})
	return config
}

type Config struct {
	Base
	Mysql
}

/*
	加载配置文件
	赋值到Config结构体
*/
func (c *Config) Load(path string) {
	//cfg,err :=ini.Load(path)
	err := ini.MapTo(c, path)
	if err != nil {
		panic(err)
	}
	//c.GinPort = cfg.Section("Base").Key("GinPort").MustInt()
	//c.GinMode = cfg.Section("Base").Key("GinMode").String()
	//c.TokenKey = cfg.Section("Base").Key("TokenKey").String()

}
