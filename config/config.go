package config

import "sync"

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
}

/*
	加载配置文件
	赋值到Config结构体
*/
func (c *Config) Load(path string) {

}
