package utils

import (
	"api/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var IMysql iMysql

type iMysql struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

func (i *iMysql) init(conf *config.Mysql) (*gorm.DB, error) {
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
	IMysql.Master = db

	return db, nil
}

func (i *iMysql) InitMaster(conf *config.Mysql) error {
	db, err := i.init(conf)
	if err != nil {
		return err
	}
	i.Master = db
	return nil
}
func (i *iMysql) InitSlave(conf *config.Mysql) error {
	db, err := i.init(conf)
	if err != nil {
		return err
	}
	i.Slave = db
	return nil
}
