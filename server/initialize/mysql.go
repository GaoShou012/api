package initialize

import (
	"api/config"
	"api/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func InitMysqlMaster(conf *config.Mysql) error {
	db, err := ConnectMysql(conf)
	global.DBMaster = db
	return err
}

func InitMysqlSlave(conf *config.Mysql) error {
	db, err := ConnectMysql(conf)
	global.DBSlave = db
	return err
}

func ConnectMysql(conf *config.Mysql) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.DNS)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(conf.PoolMax)
	db.DB().SetMaxIdleConns(conf.PoolMin)
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(conf.ConnMaxLifeTime))

	if err := db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("Ping the database=%s err=%v\n", conf.DNS, err)
	}

	db.LogMode(conf.LogMode)
	global.DBMaster = db

	return db, nil
}
