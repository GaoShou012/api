package config

func LocalLoadGaoShou() *Config {
	configInit.Do(func() {
		config = &Config{}
		config.Base = &Base{
			GinMode:  "debug",
			GinPort:  1234,
			TokenKey: "NSr9j&Z833O^iXTA",
		}

		msql := &Mysql{
			DNS:             "root:123456@tcp(127.0.0.1:13306)/bob_kf?charset=utf8mb4&loc=Local&parseTime=True",
			PoolMax:         200,
			PoolMin:         20,
			LogMode:         true,
			ConnMaxLifeTime: 600,
		}
		config.MysqlMaster = msql
		config.MysqlSlave = msql

		redisConf := &Redis{
			Addr:     "127.0.0.1:17001",
			Port:     17001,
			Password: "",
			PoolMax:  100,
			PoolMin:  10,
		}
		config.Redis = redisConf
	})

	return config
}
