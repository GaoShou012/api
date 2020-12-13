package config

import (
	"fmt"
	"path"
	"runtime"
)

func LocalLoad() *Config {
	configInit.Do(func() {
		config = &Config{}

		// config file
		{
			_, filename, _, _ := runtime.Caller(0)
			filePath := path.Join(path.Dir(filename), "dev.ini")
			fmt.Printf("配置文件路径:%s\n", filePath)
			config.Load(filePath)
		}

		{
			_, filename, _, _ := runtime.Caller(0)
			filePath := path.Join(path.Dir(filename), "ipipfree.ipdb")
			fmt.Printf("IP位置库文件路径:%s\n", filePath)
			ipLocation := &IpLocation{Path: filePath}
			config.IpLocation = ipLocation
		}
	})

	return config
}
